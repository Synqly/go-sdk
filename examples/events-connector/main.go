package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/synqly/go-sdk/client/engine"
	"github.com/synqly/go-sdk/client/engine/events"
	"github.com/synqly/go-sdk/client/management"
)

var (
	synqlyOrgId    = os.Getenv("SYNQLY_ORG_ID")
	synqlyOrgToken = os.Getenv("SYNQLY_ORG_TOKEN")
	splunkURL      = os.Getenv("SPLUNK_URL")
	splunkToken    = os.Getenv("SPLUNK_TOKEN")
)

var consoleLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

// App represents your application. We are not going to implement any app
// functionality; instead, merely keep a list of tenants, where a tenant
// represents one of your customers.
type App struct {
	Tenants []*Tenant
}

// NewApp instantiates a new App
func NewApp() *App {
	return &App{}
}

// Tenant represents one of your customers and the state you would maintain for
// them in your application. In this example, we give each tenant a unique ID
// that you would use internally to identify the tenant as well as the
// corresponding Synqly Account ID and Synqly Client that you would use to
// interact with the tenant in Synqly. Finally, we give each tenant an Event
// Logger that we will use to log events for the tenant. This uses Synqly's
// Event Connector to log events to a third-party event logging provider.
type Tenant struct {
	ID              string
	SynqlyAccountId string
	SynqlyClient    management.Client
	EventLogger     engine.Client
}

// NewTenant adds a tenant to your application. It returns an error if a tenant
// with the same ID already exists or if an account in Synqly cannot be created
// for the tenant.
func (a *App) NewTenant(ctx context.Context, id string) error {
	// Do not allow duplicate tenant names
	for _, tenant := range a.Tenants {
		if tenant.ID == id {
			return fmt.Errorf("duplicate tenant name")
		}
	}

	// Create a Synqly Account for this tenant
	account, err := management.NewClient(
		management.ClientWithBaseURL("https://api.synqly.com"),
		management.ClientWithAuthBearer(synqlyOrgToken),
		management.ClientWithHeaderSynqlyIntegrator(synqlyOrgId),
	).Accounts().CreateAccount(ctx, &management.CreateAccountRequest{
		Name: id,
	})
	if err != nil {
		return err
	}

	// Create a Synqly Client that can be used to interact with the tenant
	// We will use this client to set up an Integration with an event logging provider
	// once a user belonging to the tenant configures it. This configuration is typically
	// done in the UI, but we will do in this example to demonstrate how it works.
	client := management.NewClient(
		management.ClientWithBaseURL("https://api.synqly.com"),
		management.ClientWithHeaderSynqlyAccount(account.Result.Account.Id),
		management.ClientWithAuthBearer(account.Result.Token.Access.Secret),
	)

	a.Tenants = append(a.Tenants, &Tenant{
		ID:              id,
		SynqlyAccountId: account.Result.Account.Id,
		SynqlyClient:    client,
		EventLogger:     nil,
	})

	return nil
}

// configureEventLogging configures event logging for a tenant. It returns an
// error if the tenant cannot be found or if the Integration cannot be created
// for the tenant. This example uses Splunk as the event logging provider; you
// can use any provider that Synqly supports (see https://docs.synqly.com).
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
	credential, err := tenant.SynqlyClient.Credentials().CreateCredential(ctx, &management.Credential{
		Name: "Splunk Login",
		Type: management.CredentialTypeToken,
		Token: &management.TokenCredential{
			Secret: splunkToken,
		},
	})
	if err != nil {
		return err
	}

	// Configure an Integration for the tenant that connects to Splunk
	integration, err := tenant.SynqlyClient.Integrations().CreateIntegration(ctx, &management.CreateIntegrationRequest{
		Name:         "Background Event Logger",
		Category:     "events",
		ProviderType: "splunk",
		ProviderConfig: management.NewProviderConfigFromEvents(&management.EventConfig{
			Url:          &splunkURL,
			CredentialId: credential.Result.Id,
		}),
	})
	if err != nil {
		return err
	}

	// Create an Event Logger for the tenant
	// We will use this client to log events for the tenant
	tenant.EventLogger = engine.NewClient(
		engine.ClientWithBaseURL("https://api.synqly.com"),
		engine.ClientWithHeaderSynqlyAccount(tenant.SynqlyAccountId),
		engine.ClientWithAuthBearer(integration.Result.Token.Access.Secret),
	)

	return nil
}

// backgroundJob simulates a background job that runs for each tenant. In this
// example, we print out a message for each tenant every 10 seconds and then
// log an event if the tenant has an event logger configured. The event logging
// code does not depend on the provider configured for the tenant. Instead it
// calls the Synqly API and Synqly sends the event to provider configured for
// the tenant.
func (app *App) backgroundJob() {
	// Do some work for each tenant every 10 seconds
	for {
		for _, tenant := range app.Tenants {
			consoleLogger.Printf("Doing some important work for tenant %s\n", tenant.ID)

			// Log the result of the work to Synqly
			if tenant.EventLogger != nil {
				err := tenant.EventLogger.Events().PostEvent(context.Background(),
					// Synqly uses OCSF for events (https://ocsf.io/)
					// This creates an OCSF ScheduledJobActivity Event
					engine.NewEventFromScheduledJobActivity(&events.ScheduledJobActivity{
						ActivityId: events.ScheduledJobActivityActivityIdUpdate,
						Device: &events.Device{
							TypeId: events.DeviceTypeIdServer,
						},
						Job: &events.Job{
							File: &events.File{
								Name:   "main.go",
								TypeId: events.FileTypeIdRegularFile,
							},
							Name: "Background Job",
						},
						Metadata: &events.Metadata{
							Product: &events.Product{
								VendorName: "Synqly SDK for Go",
							},
							Version: "1.0.0",
						},
						Time:       int(time.Now().UTC().Unix()),
						SeverityId: events.SeverityIdInformational,
						TypeUid:    events.ScheduledJobActivityTypeUidScheduledJobActivityUpdate,
					}))
				if err != nil {
					consoleLogger.Printf("Error logging event for tenant %s: %s\n", tenant.ID, err)
				}
			}
		}

		time.Sleep(10 * time.Second)
	}

}

func (app *App) cleanup() {
	// Clean up the Accounts created in Synqly. In your application, you would
	// persist the Synqly account id and token to handle process restarts. We
	// are not doing that in this example, so we need to clean up the accounts
	// we created in Synqly.
	consoleLogger.Println("Cleaning up Synqly Accounts")
	ctx := context.Background()
	for _, tenant := range app.Tenants {
		// Deleting the account will delete all credentials and integrations associated with the account.
		if err := tenant.SynqlyClient.Accounts().DeleteAccount(ctx, tenant.SynqlyAccountId); err != nil {
			consoleLogger.Printf("Error deleting account %s: %s\n", tenant.SynqlyAccountId, err)
		}
	}
}

func main() {
	ctx := context.Background()

	if synqlyOrgId == "" || synqlyOrgToken == "" || splunkURL == "" || splunkToken == "" {
		log.Fatal("SYNQLY_ORG_ID, SYNQLY_ORG_TOKEN, SPLUNK_URL, and SPLUNK_TOKEN environment variables must be set")
	}

	// Create a couple of tenants
	app := NewApp()

	// Be sure to clean up the Synqly Accounts we create before exiting
	defer app.cleanup()

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

	go app.backgroundJob()

	// Wait for user to control c to exit
	select {}
}
