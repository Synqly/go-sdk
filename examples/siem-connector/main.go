package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"

	boff "github.com/cenkalti/backoff/v4"
	"github.com/spf13/viper"

	"github.com/synqly/go-sdk/client/engine"
	engineClient "github.com/synqly/go-sdk/client/engine/client"
	scheduledJobActivity "github.com/synqly/go-sdk/client/engine/ocsf/v110/scheduledjobactivity"
	mgmt "github.com/synqly/go-sdk/client/management"
	mgmtClient "github.com/synqly/go-sdk/client/management/client"
)

var consoleLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

// Config model of the providers for the application
type provider interface {
	ProviderConfig(credentialId string) *mgmt.ProviderConfig
	Validate() error
	configureIntegration(ctx context.Context, tenantID string, app *App) error
}

// Config model of the application
type synqlyConfig struct {
	SynqlyOrgToken  string `mapstructure:"SYNQLY_ORG_TOKEN"`
	DurationSeconds string `mapstructure:"DURATION_SECONDS"`
}

// Config splunk model of the providers for the application
type splunkConfig struct {
	SplunkHecUrl      string `mapstructure:"SPLUNK_HEC_URL"`
	SplunkHecToken    string `mapstructure:"SPLUNK_HEC_TOKEN"`
	SplunkSearchUrl   string `mapstructure:"SPLUNK_SEARCH_URL"`
	SplunkSearchToken string `mapstructure:"SPLUNK_SEARCH_TOKEN"`
}

// Config Sumo Logic model of the providers for the application
type sumoConfig struct {
	SumoLogicCollectionUrl string `mapstructure:"SUMO_LOGIC_COLLECTION_URL"`
	SumoLogicAccessId      string `mapstructure:"SUMO_LOGIC_ACCESS_ID"`
	SumoLogicAccessKey     string `mapstructure:"SUMO_LOGIC_ACCESS_KEY"`
	SumoLogicUrl           string `mapstructure:"SUMO_LOGIC_URL"`
}

type mockConfig struct {
}

func (c *mockConfig) ProviderConfig(credentialId string) *mgmt.ProviderConfig {
	return &mgmt.ProviderConfig{SiemMockSiem: &mgmt.SiemMock{}}
}

func (c *mockConfig) Validate() error {
	return nil
}

func (c *mockConfig) configureIntegration(ctx context.Context, tenantID string, app *App) error {
	// Find the tenant
	var tenant *Tenant
	for _, t := range app.Tenants {
		if t.ID == tenantID {
			tenant = t
			break
		}
	}
	if tenant == nil {
		return fmt.Errorf("tenant not found")
	}

	integration, err := tenant.SynqlyClient.Integrations.Create(ctx, tenant.SynqlyAccountId, &mgmt.CreateIntegrationRequest{
		Fullname:       mgmt.String("Event Logger and Query"),
		ProviderConfig: c.ProviderConfig(""),
	})
	if err != nil {
		return fmt.Errorf("unable to create integration: %w", err)
	}

	tenant.SynqlyEngineClient = engineClient.NewClient(
		engineClient.WithToken(integration.Result.Token.Access.Secret),
	)
	return nil
}

// ProviderConfig returns the ProviderConfig for the given provider
func (c *splunkConfig) ProviderConfig(credentialId string) *mgmt.ProviderConfig {
	config := &mgmt.ProviderConfig{
		SiemSplunk: &mgmt.SiemSplunk{
			HecUrl: c.SplunkHecUrl,
			HecCredential: &mgmt.SplunkHecToken{
				TokenId: credentialId,
			},
			// Do not verify the Splunk server's TLS certificate. This
			// is not recommended for production use; however, it is set
			// here because Splunk Cloud HEC endpoints use self-signed
			// "SplunkServerDefaultCert" certificates by default.
			SkipTlsVerify: true,
		}}

	config.SiemSplunk.SearchServiceUrl = &c.SplunkSearchUrl
	config.SiemSplunk.SearchServiceCredential = &mgmt.SplunkSearchCredential{
		Token: &mgmt.TokenCredential{
			Secret: c.SplunkSearchToken,
		},
	}
	return config
}

// Validate checks if the required fields are set in the config file for the provider
func (c *splunkConfig) Validate() error {
	if c.SplunkHecUrl == "" || c.SplunkHecToken == "" || c.SplunkSearchUrl == "" || c.SplunkSearchToken == "" {
		return errors.New("missing required Splunk configuration")
	}
	return nil
}

// configureIntegration initializes event logging for a Tenant.
func (c *splunkConfig) configureIntegration(ctx context.Context, tenantID string, app *App) error {
	// Find the tenant
	var tenant *Tenant
	for _, t := range app.Tenants {
		if t.ID == tenantID {
			tenant = t
			break
		}
	}
	if tenant == nil {
		return fmt.Errorf("tenant not found")
	}

	// Configure an Integration in Synqly depending on this tenant's config
	credential, err := tenant.SynqlyClient.Credentials.Create(ctx, tenant.SynqlyAccountId, &mgmt.CreateCredentialRequest{
		Fullname: mgmt.String(fmt.Sprintf("%s authentication token", "splunk")),
		Config: &mgmt.CredentialConfig{
			Token: &mgmt.TokenCredential{
				Secret: c.SplunkHecToken,
			},
		},
	})
	if err != nil {
		return err
	}

	providerConfig := c.ProviderConfig(credential.Result.Id)

	integration, err := tenant.SynqlyClient.Integrations.Create(ctx, tenant.SynqlyAccountId, &mgmt.CreateIntegrationRequest{
		Fullname:       mgmt.String("Event Logger and Query"),
		ProviderConfig: providerConfig,
	})
	if err != nil {
		return fmt.Errorf("unable to create integration: %w", err)
	}

	tenant.SynqlyEngineClient = engineClient.NewClient(
		engineClient.WithToken(integration.Result.Token.Access.Secret),
	)

	return nil
}

// ProviderConfig returns the ProviderConfig for the given provider
func (c *sumoConfig) ProviderConfig(credentialId string) *mgmt.ProviderConfig {
	return &mgmt.ProviderConfig{
		SiemSumoLogic: &mgmt.SiemSumoLogic{
			Credential: &mgmt.SumoLogicCredential{
				Type:    "basic",
				BasicId: c.SumoLogicAccessKey,
				Basic: &mgmt.BasicCredential{
					Username: c.SumoLogicAccessId,
					Secret:   c.SumoLogicAccessKey},
			},
			Url: c.SumoLogicUrl,
			CollectionUrl: &mgmt.SumoLogicCollectionUrl{
				Secret: &mgmt.SecretCredential{Secret: c.SumoLogicCollectionUrl},
				Type:   "secret",
			},
		},
	}
}

func (c *sumoConfig) configureIntegration(ctx context.Context, tenantID string, app *App) error {
	// Find the tenant
	var tenant *Tenant
	for _, t := range app.Tenants {
		if t.ID == tenantID {
			tenant = t
			break
		}
	}

	if tenant == nil {
		return fmt.Errorf("tenant not found")
	}

	providerConfig := c.ProviderConfig("")

	integration, err := tenant.SynqlyClient.Integrations.Create(ctx, tenant.SynqlyAccountId, &mgmt.CreateIntegrationRequest{
		Fullname:       mgmt.String("Event Logger and Query"),
		ProviderConfig: providerConfig,
	})
	if err != nil {
		return fmt.Errorf("unable to create integration: %w", err)
	}

	tenant.SynqlyEngineClient = engineClient.NewClient(
		engineClient.WithToken(integration.Result.Token.Access.Secret),
	)

	return nil
}

func (c *sumoConfig) Validate() error {
	if c.SumoLogicCollectionUrl == "" || c.SumoLogicAccessId == "" || c.SumoLogicAccessKey == "" || c.SumoLogicUrl == "" {
		return errors.New("missing required Sumo Logic configuration")
	}
	return nil
}

// App represents your application. We are not going to implement any app
// functionality; instead, merely keep a list of tenants, where a tenant
// represents an end user of your product.
type App struct {
	Tenants []*Tenant
}

// NewApp instantiates a new App
func NewApp() *App {
	return &App{}
}

// A Tenant object represents one of your customers, as well as the state you
// would maintain for them in your application.
type Tenant struct {
	// ID: A unique identifier for one of your customers.
	ID string
	// SynqlyAccountId: The ID of the Synqly Account in which Integrations for
	// this tenant will be created. This example creates a new Account for every
	// Tenant, but it would also be possible to use the same Account for all Tenants.
	SynqlyAccountId string
	// SynqlyClient: A cached management client, used to manage Integrations.
	SynqlyClient *mgmtClient.Client
	// SynqlyEngineClient: A cached engine client, used to log events to a third-party
	// logging Provider by way of an Integration.
	SynqlyEngineClient *engineClient.Client
}

// NewTenant initializes a Synqly Account for a given Tenant. This example
// creates a new Account for every Tenant to keep their credentials isolated.
//
// Returns an error if a tenant with the same ID already exists or if a Synqly
// Account cannot be created for the Tenant.
func (app *App) NewTenant(ctx context.Context, id string, synqlyOrgToken string) error {
	// Do not allow duplicate tenant names
	for _, tenant := range app.Tenants {
		if tenant.ID == id {
			return fmt.Errorf("duplicate tenant name %v", id)
		}
	}

	// Create a Synqly Client that can be used to interact with the tenant
	// We will use this client to create an Account and set up an Integration with an event logging provider
	client := mgmtClient.NewClient(
		mgmtClient.WithToken(synqlyOrgToken),
	)

	// Create a Synqly Account for this tenant
	account, err := client.Accounts.Create(ctx, &mgmt.CreateAccountRequest{
		Fullname: &id,
	})
	if err != nil {
		return fmt.Errorf("unable to create account: %w", err)
	}

	// Store the Tenant and associated Synqly objects in an in-memory cache.
	app.Tenants = append(app.Tenants, &Tenant{
		ID:                 id,
		SynqlyAccountId:    account.Result.Account.Id,
		SynqlyClient:       client,
		SynqlyEngineClient: nil,
	})
	return nil
}

// query performs a query against the SIEM provider for each tenant.
func (app *App) query(ctx context.Context) error {
	req := &engine.QuerySiemEventsRequest{
		Cursor:         nil,
		IncludeRawData: engine.Bool(true),
		Order:          []*string{engine.String("time[asc]")},
		Limit:          engine.Int(10),
	}

	for _, tenant := range app.Tenants {
		res, err := boff.RetryWithData(func() (*engine.QuerySiemEventsResponse, error) {
			res, err := tenant.SynqlyEngineClient.Siem.QueryEvents(ctx, req)
			if err != nil {
				return res, boff.Permanent(err)
			}
			if res != nil && res.Status == engine.QueryStatusComplete {
				return res, nil
			}
			if res.Status == engine.QueryStatusPending {
				req.Cursor = &res.Cursor
				return nil, fmt.Errorf("query is pending")
			}
			return res, boff.Permanent(fmt.Errorf("unexpected query status: %s", res.Status))
		}, boff.WithMaxRetries(&boff.ZeroBackOff{}, 10))

		if err != nil {
			return fmt.Errorf("unable to query events: %w", err)
		}

		consoleLogger.Println("Printing some fields from results")
		for _, event := range res.Result {
			metadata, ok := event["metadata"].(map[string]interface{})
			if !ok {
				consoleLogger.Println("metadata is nil or not a map")
				continue
			}
			job, ok := event["job"].(map[string]interface{})
			if !ok {
				consoleLogger.Println("job is nil or not a map")
				continue
			}
			consoleLogger.Printf("job: %v\n", job["name"])
			consoleLogger.Printf("log_name: %v\n", metadata["log_name"])
			consoleLogger.Printf("class_name: %v\n", event["class_name"])
			consoleLogger.Printf("type_uid: %v\n", event["type_uid"])
		}
		consoleLogger.Println()
	}
	return nil
}
func (app *App) splunkConfig(splunkHecURL, splunkSearchURL, splunkHecToken, splunkSearchToken string) *mgmt.ProviderConfig {
	config := &mgmt.ProviderConfig{
		SiemSplunk: &mgmt.SiemSplunk{
			HecUrl: splunkHecURL,
			HecCredential: &mgmt.SplunkHecToken{
				TokenId: splunkHecToken,
			},
			// Do not verify the Splunk server's TLS certificate. This
			// is not recommended for production use; however, it is set
			// here because Splunk Cloud HEC endpoints use self-signed
			// "SplunkServerDefaultCert" certificates by default.
			SkipTlsVerify: true,
		}}

	if splunkSearchToken != "" {
		config.SiemSplunk.SearchServiceUrl = &splunkSearchURL
		config.SiemSplunk.SearchServiceCredential = &mgmt.SplunkSearchCredential{
			Token: &mgmt.TokenCredential{
				Secret: splunkSearchToken,
			},
		}
	}
	return config
}

func (app *App) inmemConfig() *mgmt.ProviderConfig {
	return &mgmt.ProviderConfig{SiemMockSiem: &mgmt.SiemMock{}}
}

// example, we print a message for every Tenant tenant at regular intervals,
// backgroundJob simulates a background job that runs for every Tenant. In this
// and log the message as an event if the Tenant's event logger is configured.
//
// Note that the PostEvent() call is not specific to Splunk. The code would be
// the same for any supported Event Provider. This is where Synqly's abstraction
// kicks in, Integrations within a given Connector category (in this case, Events)
// share a unified API for data operations, no matter which Provider they target.
func (app *App) backgroundJob(durationSeconds int) {
	ctx := context.Background()

	startTime := time.Now()
	endTime := startTime.Add(time.Duration(durationSeconds) * time.Second)

	// Loop infinitely, generating data at 1 second intervals
	for {
		for _, tenant := range app.Tenants {
			consoleLogger.Printf("Doing some work for tenant %s\n", tenant.ID)

			// Call a util function to generate a sample OCSF Event.
			newEvent := createSampleEvent()

			if time.Now().UTC().After(endTime) {
				consoleLogger.Println("Done generating events")
				return
			}

			// If the EventLogger for the given Tenant has been initialized, use
			// it to post data.
			if tenant.SynqlyEngineClient != nil {
				// Log the result of the work to Synqly
				err := tenant.SynqlyEngineClient.Siem.PostEvents(
					ctx,
					[]*engine.Event{newEvent},
				)
				if err != nil {
					consoleLogger.Printf("error logging event for tenant %s: %s\n", tenant.ID, err)
				}
				consoleLogger.Printf("Logged event for tenant %s\n", tenant.ID)
			}
		}
		time.Sleep(1*time.Second + time.Duration(rand.Intn(1000))*time.Millisecond) //nolint: gosec
		fmt.Println()
	}
}

// createSampleEvent generates a sample ScheduledJobActivity OCSF (https://ocsf.io/) Event
func createSampleEvent() *engine.Event {
	actionId := scheduledJobActivity.Action_Allowed
	return &engine.Event{
		ScheduledJobActivity: &scheduledJobActivity.ScheduledJobActivity{
			ActivityId: scheduledJobActivity.Activity_Update,
			ActionId:   &actionId,
			Device: &scheduledJobActivity.Device{
				TypeId: scheduledJobActivity.Device_Type_Server,
			},
			Job: &scheduledJobActivity.Job{
				File: &scheduledJobActivity.File{
					Name:   "main.go",
					TypeId: 1,
				},
				Name: "Background Job",
			},
			Metadata: &scheduledJobActivity.Metadata{
				Product: &scheduledJobActivity.Product{
					VendorName: "Synqly SDK for Go",
				},
				Version: "1.1.0",
			},
			Time:       int(time.Now().UTC().Unix()),
			SeverityId: scheduledJobActivity.Severity_Informational,
			TypeUid:    scheduledJobActivity.Type_ScheduledJobActivity_Update,
		},
	}
}

func (app *App) cleanup() {
	// Clean up the Accounts created in Synqly. In your application, you would
	// persist the Synqly Account id and Integration tokens to handle process
	// restarts. We are not doing that in this example, so we need to clean up
	// the accounts we created in Synqly.
	consoleLogger.Println("Cleaning up Synqly Accounts")
	ctx := context.Background()
	for _, tenant := range app.Tenants {
		// Deleting the account will delete all credentials and integrations associated with the account.
		if err := tenant.SynqlyClient.Accounts.Delete(ctx, tenant.SynqlyAccountId); err != nil {
			consoleLogger.Printf("Error deleting account %s: %s\n", tenant.SynqlyAccountId, err)
		}
	}
	os.Exit(0)
}

func main() {
	ctx := context.Background()

	synqlyConfig := synqlyConfig{}
	splunkConfig := splunkConfig{}
	sumoConfig := sumoConfig{}
	mockConfig := mockConfig{}

	// Load config variables from the env file
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&synqlyConfig)
	if err != nil {
		fmt.Println("Error unmarshalling synqlyConfig:", err)
		return
	}

	err = viper.Unmarshal(&splunkConfig)
	if err != nil {
		fmt.Println("Error unmarshalling splunkConfig:", err)
		return
	}

	err = viper.Unmarshal(&sumoConfig)
	if err != nil {
		fmt.Println("Error unmarshalling sumoConfig:", err)
		return
	}

	// Check for required environment variables
	if synqlyConfig.SynqlyOrgToken == "" {
		log.Fatal("Must set following environment variable: SYNQLY_ORG_TOKEN")
	}
	err = splunkConfig.Validate()
	if err != nil {
		consoleLogger.Print("WARNING: no Splunk credentials provided (SLUNK_URL, SPLUNK_HEC_TOKEN)\nUsing Mock as the SIEM  provider")
	}
	err = sumoConfig.Validate()
	if err != nil {
		consoleLogger.Print("WARNING: no Sumo credentials provided")
	}

	// Instantiate App object
	app := NewApp()

	// Create an interrupt handler to clean up tenants if the program is shut down
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		// Intercept all ^C
		for range c {
			app.cleanup()
		}
	}()
	// Also be sure to run clean up if the program exits gracefully
	defer app.cleanup()

	// Create and configure Tenant ABC to use splunk SIEM event logging provider
	consoleLogger.Print("Creating Tenant ABC with splunk SIEM provider")
	if err := app.NewTenant(ctx, "Tenant ABC", synqlyConfig.SynqlyOrgToken); err != nil {
		log.Fatal(err)
	}
	if err := splunkConfig.configureIntegration(ctx, "Tenant ABC", app); err != nil {
		log.Fatal(err)
	}

	// Create and configure Tenant DEF to use sumo SIEM event logging provider
	consoleLogger.Print("Creating Tenant DEF with sumo SIEM provider")
	if err := app.NewTenant(ctx, "Tenant DEF", synqlyConfig.SynqlyOrgToken); err != nil {
		log.Fatal(err)
	}

	if err := sumoConfig.configureIntegration(ctx, "Tenant DEF", app); err != nil {
		fmt.Println(err)
		app.cleanup()
		log.Fatal(err)
	}

	// Create and configure Tenant XYZ to use mock SIEM event logging provider
	consoleLogger.Print("Creating Tenant XYZ with mock SIEM provider")
	if err := app.NewTenant(ctx, "Tenant XYZ", synqlyConfig.SynqlyOrgToken); err != nil {
		log.Fatal(err)
	}
	if err := mockConfig.configureIntegration(ctx, "Tenant XYZ", app); err != nil {
		log.Fatal(err)
	}

	// Query the SIEM provider for each tenant
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := app.query(ctx); err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	// Generate synthetic load for the tenants
	if synqlyConfig.DurationSeconds == "" {
		// If no duration provided, run for 10m
		app.backgroundJob(600)
	} else {
		// Otherwise, run for the provided duration
		dur, err := strconv.Atoi(synqlyConfig.DurationSeconds)
		if err != nil {
			log.Fatal(err)
		}
		app.backgroundJob(dur)
	}
}
