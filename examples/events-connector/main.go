package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	engine "github.com/synqly/go-sdk/client/engine"
	engineClient "github.com/synqly/go-sdk/client/engine/client"

	// Each OSCF event type has its own package. This is intended to make imports
	// more granular, allowing the end-user to import only the types they need.
	scheduledJobActivity "github.com/synqly/go-sdk/client/engine/ocsf/scheduledjobactivity"

	mgmt "github.com/synqly/go-sdk/client/management"
	mgmtClient "github.com/synqly/go-sdk/client/management/client"
)

var (
	// SYNQLY_ORG_ID: Your Synqly Organization ID
	synqlyOrgId = os.Getenv("SYNQLY_ORG_ID")
	// SYNQLY_ORG_TOKEN: A Synqly Organization Token, used to create Accounts and
	// Integrations
	synqlyOrgToken = os.Getenv("SYNQLY_ORG_TOKEN")
	// SPLUNK_URL: URL of a Splunk HTTP event collector endpoint
	// Example: "https://prd-p-icwnd.splunkcloud.com:8088/services/collector/event"
	splunkURL = os.Getenv("SPLUNK_URL")
	// SPLUNK_HEC_TOKEN: A Splunk HTTP Event Collector token for logging events
	splunkToken = os.Getenv("SPLUNK_HEC_TOKEN")
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
	// SynqlyClient: A cached managmenet client, used to manage Integrations.
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
			return fmt.Errorf("duplicate tenant name")
		}
	}

	// Create a Synqly Account for this tenant
	account, err := mgmtClient.NewClient(
		mgmtClient.WithAuthToken(synqlyOrgToken),
	).Accounts.CreateAccount(ctx, &mgmt.CreateAccountRequest{
		Name: id,
	})
	if err != nil {
		return fmt.Errorf("unable to create account: %w", err)
	}

	// Create a Synqly Client that can be used to interact with the tenant
	// We will use this client to set up an Integration with an event logging provider
	// once a user belonging to the tenant configures it. This configuration is typically
	// done in the UI, but we will do in this example to demonstrate how it works.
	client := mgmtClient.NewClient(
		mgmtClient.WithAuthToken(synqlyOrgToken),
	)

	// Store the Tenant and associated Synqly objects in an in-memory cache.
	a.Tenants = append(a.Tenants, &Tenant{
		ID:              id,
		SynqlyAccountId: account.Result.Account.Id,
		SynqlyClient:    client,
		EventLogger:     nil,
	})

	return nil
}

// configureEventLogging initializes event logging for a Tenant. Stores the
// HEC_TOKEN as a secure Credential object, then creates a Splunk Integration
// targeting SPLUNK_URL. This example uses Splunk as the event logging
// provider; however, an Integration can be configured to target any supported
// Provider type.
//
// Returns an error if the Tenant cannot be found, or if an Integration cannot
// be created for the given Tenant.
func (a *App) configureEventLogging(ctx context.Context, tenantID string) error {
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

	// We need to save the tenant's Splunk credentials in Synqly before configuring the Integration
	// We will use the Synqly Client we created for the tenant to do this
	credential, err := tenant.SynqlyClient.Credentials.CreateCredential(ctx, tenant.SynqlyAccountId, &mgmt.CreateCredentialRequest{
		Name: "Splunk Login",
		Config: mgmt.NewCredentialConfigFromToken(&mgmt.TokenCredential{
			Secret: splunkToken,
		}),
	})
	if err != nil {
		return err
	}

	// Configure an Integration for the tenant that connects to Splunk
	integration, err := tenant.SynqlyClient.Integrations.CreateIntegration(ctx, tenant.SynqlyAccountId, &mgmt.CreateIntegrationRequest{
		Name:         "Background Event Logger",
		Category:     "events",
		ProviderType: "splunk",
		// Set Splunk Config to ignore cert checking on the Splunk endpoint.
		// This is not recommended for production use, we only set it here
		// because Splunk HEC endpoints use self-signed "SplunkServerDefaultCert"
		// certificates when first initialized.
		ProviderConfig: mgmt.NewProviderConfigFromEvents(&mgmt.EventConfig{
			Url:          &splunkURL,
			CredentialId: credential.Result.Id,
			Config: &mgmt.EventProviderTypeConfig{
				Type: "splunk",
				Splunk: &mgmt.SplunkConfig{
					SkipTlsVerify: true,
				},
			},
		}),
	})
	if err != nil {
		return fmt.Errorf("unable to create integration: %w", err)
	}

	// Create an Event Logger for the tenant
	// We will use this client to log events for the tenant
	tenant.EventLogger = engineClient.NewClient(
		engineClient.WithAuthToken(integration.Result.Token.Access.Secret),
	)

	return nil
}

// backgroundJob simulates a background job that runs for every Tenant. In this
// example, we print a message for every Tenant tenant at regular intervals,
// and log the message as an event if the Tenant's event logger is configured.
//
// Note that the PostEvent() call is not specific to Splunk. The code would be
// the same for any supported Event Provider. This is where Synqly's abstraction
// kicks in, Integrations within a given Connector category (in this case, Events)
// share a unified API for data operations, no matter which Provider they target.
func (app *App) backgroundJob() {
	// Loop infinitely, generating data at 1 second intervals
	for {
		for _, tenant := range app.Tenants {
			consoleLogger.Printf("Doing some work for tenant %s\n", tenant.ID)

			// Call a helper function to generate a sample OCSF Event.
			newEvent := createSampleEvent()

			// If the EventLogger for the given Tenant has been initialized, use
			// it to send data.
			if tenant.EventLogger != nil {
				// Log the result of the work to Synqly
				err := tenant.EventLogger.Events.PostEvent(
					context.Background(),
					newEvent,
				)
				if err != nil {
					consoleLogger.Printf("error logging event for tenant %s: %s\n",
						tenant.ID, err)
				}
			}
		}

		time.Sleep(1 * time.Second)
	}
}

// createSampleEvent generates a sample ScheduledJobActivity OCSF (https://ocsf.io/) Event
func createSampleEvent() *engine.Event {
	return engine.NewEventFromScheduledJobActivity(&scheduledJobActivity.ScheduledJobActivity{
		ActivityId: scheduledJobActivity.Activity_Update,
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
			Version: "1.0.0",
		},
		Time:       int(time.Now().UTC().Unix()),
		SeverityId: scheduledJobActivity.Severity_Informational,
		TypeUid:    scheduledJobActivity.Type_ScheduledJobActivity_Update,
	})
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
		if err := tenant.SynqlyClient.Accounts.DeleteAccount(ctx, tenant.SynqlyAccountId); err != nil {
			consoleLogger.Printf("Error deleting account %s: %s\n", tenant.SynqlyAccountId, err)
		}
	}
	os.Exit(0)
}

func main() {
	ctx := context.Background()

	if synqlyOrgId == "" || synqlyOrgToken == "" || splunkURL == "" || splunkToken == "" {
		log.Fatal("Must set following environment variables: SYNQLY_ORG_ID, SYNQLY_ORG_TOKEN, SPLUNK_URL, SPLUNK_HEC_TOKEN")
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
	if err := app.NewTenant(ctx, "Tenant ABC"); err != nil {
		log.Fatal(err)
	}
	if err := app.NewTenant(ctx, "Tenant XYZ"); err != nil {
		log.Fatal(err)
	}

	// Configure one tenant to use event logging
	if err := app.configureEventLogging(ctx, "Tenant ABC"); err != nil {
		log.Fatal(err)
	}

	// Generate synthetic load for the tenants
	go app.backgroundJob()

	// Wait for user to control c to exit
	select {}
}
