package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"

	engine "github.com/synqly/go-sdk/client/engine"
	engineClient "github.com/synqly/go-sdk/client/engine/client"

	// Each OCSF event type has its own package. This is intended to make imports
	// more granular, allowing the end-user to import only the types they need.
	scheduledJobActivity "github.com/synqly/go-sdk/client/engine/ocsf/v110/scheduledjobactivity"

	mgmt "github.com/synqly/go-sdk/client/management"
	mgmtClient "github.com/synqly/go-sdk/client/management/client"
)

var (
	// SYNQLY_ORG_TOKEN: A Synqly Organization Token, used to create Accounts and
	// Integrations
	synqlyOrgToken = os.Getenv("SYNQLY_ORG_TOKEN")
	// SPLUNK_URL: URL of a Splunk HTTP event collector endpoint
	// Example: "https://prd-p-icwnd.splunkcloud.com:8088/services/collector/event"
	splunkURL = os.Getenv("SPLUNK_URL")
	// SPLUNK_HEC_TOKEN: A Splunk HTTP Event Collector token for logging events
	splunkToken = os.Getenv("SPLUNK_HEC_TOKEN")
	// DURATION_SECONDS: (Optional) limits event generation to the provided duration
	durationSeconds = os.Getenv("DURATION_SECONDS")
)

var consoleLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

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
	// EventLogger: A cached engine client, used to log events to a third-party
	// logging Provider by way of an Integration.
	EventLogger *engineClient.Client
}

// NewTenant initializes a Synqly Account for a given Tenant. This example
// creates a new Account for every Tenant to keep their credentials isolated.
//
// Returns an error if a tenant with the same ID already exists or if a Synqly
// Account cannot be created for the Tenant.
func (a *App) NewTenant(ctx context.Context, id string) error {
	// Do not allow duplicate tenant names
	for _, tenant := range a.Tenants {
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
	a.Tenants = append(a.Tenants, &Tenant{
		ID:              id,
		SynqlyAccountId: account.Result.Account.Id,
		SynqlyClient:    client,
		EventLogger:     nil,
	})
	return nil
}

// configureEventLogging initializes event logging for a Tenant.
// This example uses Splunk SIEM providers and mock SIEM provider as the event logging
// providers; however, an Integration can be configured to target any supported provider type.
// Stores the HEC_TOKEN as a secure Credential object, then creates a Splunk Integration
// targeting SPLUNK_URL.
//
// Returns an error if the Tenant cannot be found, or if an Integration cannot
// be created for the given Tenant.
func (a *App) configureEventLogging(ctx context.Context, tenantID, siemProviderType string) error {
	// Find the tenant
	var tenant *Tenant
	for _, t := range a.Tenants {
		if t.ID == tenantID {
			tenant = t
			break
		}
	}
	if tenant == nil {
		return fmt.Errorf("tenant not found")
	}

	// Configure an Integration in Synqly depending on this tenant's config
	var providerConfig *mgmt.ProviderConfig
	switch siemProviderType {
	case "splunk":
		// We need to save the tenant's Splunk credentials in Synqly before configuring the Integration
		// We will use the Synqly Client we created for the tenant to do this
		credential, err := tenant.SynqlyClient.Credentials.Create(ctx, tenant.SynqlyAccountId, &mgmt.CreateCredentialRequest{
			Fullname: mgmt.String(fmt.Sprintf("%s authentication token", siemProviderType)),
			Config: &mgmt.CredentialConfig{
				Token: &mgmt.TokenCredential{
					Secret: splunkToken,
				},
			},
		})
		if err != nil {
			return err
		}

		providerConfig = a.splunkConfig(splunkURL, credential.Result.Id)

	case "inmem":
		providerConfig = a.inmemConfig()

	default:
		return fmt.Errorf("invalid siem provider type: %s", siemProviderType)
	}

	integration, err := tenant.SynqlyClient.Integrations.Create(ctx, tenant.SynqlyAccountId, &mgmt.CreateIntegrationRequest{
		Fullname:       mgmt.String("Background Event Logger"),
		ProviderConfig: providerConfig,
	})
	if err != nil {
		return fmt.Errorf("unable to create integration: %w", err)
	}

	// Create an Event Logger for the tenant
	// We will use this client to post events to the provider
	tenant.EventLogger = engineClient.NewClient(
		engineClient.WithToken(integration.Result.Token.Access.Secret),
	)

	return nil
}

func (a *App) splunkConfig(splunkURL, credentialId string) *mgmt.ProviderConfig {
	return &mgmt.ProviderConfig{
		SiemSplunk: &mgmt.SiemSplunk{
			HecUrl: splunkURL,
			HecCredential: &mgmt.SplunkHecToken{
				TokenId: credentialId,
			},
			// Do not verify the Splunk server's TLS certificate. This
			// is not recommended for production use; however, it is set
			// here because Splunk Cloud HEC endpoints use self-signed
			// "SplunkServerDefaultCert" certificates by default.
			SkipTlsVerify: true,
		}}
}

func (a *App) inmemConfig() *mgmt.ProviderConfig {
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

			// Call a helper function to generate a sample OCSF Event.
			newEvent := createSampleEvent()

			if time.Now().UTC().After(endTime) {
				consoleLogger.Println("Done generating events")
				return
			}

			// If the EventLogger for the given Tenant has been initialized, use
			// it to post data.
			if tenant.EventLogger != nil {
				// Log the result of the work to Synqly
				err := tenant.EventLogger.Siem.PostEvents(
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
	return &engine.Event{
		ScheduledJobActivity: &scheduledJobActivity.ScheduledJobActivity{
			ActivityId: scheduledJobActivity.Activity_Update,
			ActionId:   scheduledJobActivity.Action_Allowed,
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

	if synqlyOrgToken == "" {
		log.Fatal("Must set following environment variable: SYNQLY_ORG_TOKEN")
	}
	if splunkURL == "" || splunkToken == "" {
		consoleLogger.Print("WARNING: no Splunk credentials provided (SLUNK_URL, SPLUNK_HEC_TOKEN)\nUsing Mock as the SIEM  provider")
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

	// Create a couple of tenants

	if splunkToken != "" && splunkURL != "" {
		// Create and configure Tenant ABC to use splunk SIEM event logging provider
		consoleLogger.Print("Creating Tenant ABC with splunk SIEM provider")
		if err := app.NewTenant(ctx, "Tenant ABC"); err != nil {
			log.Fatal(err)
		}
		if err := app.configureEventLogging(ctx, "Tenant ABC", "splunk"); err != nil {
			log.Fatal(err)
		}
	}

	// Create and configure Tenant XYZ to use mock SIEM event logging provider
	consoleLogger.Print("Creating Tenant XYZ with mock SIEM provider")
	if err := app.NewTenant(ctx, "Tenant XYZ"); err != nil {
		log.Fatal(err)
	}
	if err := app.configureEventLogging(ctx, "Tenant XYZ", "inmem"); err != nil {
		log.Fatal(err)
	}

	// Generate synthetic load for the tenants
	if durationSeconds == "" {
		// If no duration provided, run for 10m
		app.backgroundJob(600)
	} else {
		// Otherwise, run for the provided duration
		dur, err := strconv.Atoi(durationSeconds)
		if err != nil {
			log.Fatal(err)
		}
		app.backgroundJob(dur)
	}
}
