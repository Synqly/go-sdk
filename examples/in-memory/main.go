package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	engine "github.com/synqly/go-sdk/client/engine"
	engineClient "github.com/synqly/go-sdk/client/engine/client"
	scheduledJobActivity "github.com/synqly/go-sdk/client/engine/ocsf/scheduledjobactivity"

	mgmt "github.com/synqly/go-sdk/client/management"
	mgmtClient "github.com/synqly/go-sdk/client/management/client"
)

var (
	synqlyOrgId    = os.Getenv("SYNQLY_ORG_ID")
	synqlyOrgToken = os.Getenv("SYNQLY_ORG_TOKEN")
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
	SynqlyClient    mgmtClient.Client
	EventLogger     *engineClient.Client
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
	account, err := mgmtClient.NewClient(
		mgmtClient.WithBaseURL("https://api.synqly.com"),
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
		mgmtClient.WithBaseURL("https://api.synqly.com"),
		mgmtClient.WithAuthToken(synqlyOrgToken),
	)

	a.Tenants = append(a.Tenants, &Tenant{
		ID:              id,
		SynqlyAccountId: account.Result.Account.Id,
		SynqlyClient:    *client,
		// TODO: Check if we can set this to nil
		EventLogger: nil,
	})
	log.Printf("Account %v created for tenant %v", account.Result.Account.Id, id)

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
		return fmt.Errorf("tenant %v not found", tenantID)
	} else {
		consoleLogger.Printf("Found Synqly account with account ID %v for tenant %v", tenant.SynqlyAccountId, tenantID)
	}

	// We need to save the tenant's Splunk credentials in Synqly before configuring the Integration
	// We will use the Synqly Client we created for the tenant to do this
	credential, err := tenant.SynqlyClient.Credentials.CreateCredential(ctx, tenant.SynqlyAccountId, &mgmt.Credential{
		Name: "Inmem Login",
		Type: mgmt.CredentialTypeToken,
		Token: &mgmt.TokenCredential{
			Secret: "in-mem-sample-cred",
		},
	})
	if err != nil {
		return err
	}

	// Configure an Integration for the tenant that connects to Splunk
	integration, err := tenant.SynqlyClient.Integrations.CreateIntegration(ctx, tenant.SynqlyAccountId, &mgmt.CreateIntegrationRequest{
		Name:         "inmem-demo-provider",
		Category:     "events",
		ProviderType: "inmem",
		ProviderConfig: mgmt.NewProviderConfigFromEvents(&mgmt.EventConfig{
			Url:          nil,
			CredentialId: credential.Result.Id,
		}),
	})
	if err != nil {
		return fmt.Errorf("unable to create integration: %w", err)
	}

	// Create an Event Logger for the tenant
	// We will use this client to log events for the tenant
	tenant.EventLogger = engineClient.NewClient(
		engineClient.WithBaseURL("https://api.synqly.com"),
		engineClient.WithAuthToken(integration.Result.Token.Access.Secret),
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

			if tenant.EventLogger != nil {
				// Log the result of the work to Synqly
				err := tenant.EventLogger.Events.PostEvent(context.Background(),
					// Synqly uses OCSF for events (https://ocsf.io/)
					// This creates an OCSF ScheduledJobActivity Event
					engine.NewEventFromScheduledJobActivity(&scheduledJobActivity.ScheduledJobActivity{
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
		if err := tenant.SynqlyClient.Accounts.DeleteAccount(ctx, tenant.SynqlyAccountId); err != nil {
			consoleLogger.Printf("Error deleting account %s: %s\n", tenant.SynqlyAccountId, err)
		}
	}
}

func main() {
	ctx := context.Background()

	if synqlyOrgId == "" || synqlyOrgToken == "" {
		log.Fatal("SYNQLY_ORG_ID and SYNQLY_ORG_TOKEN environment variables must be set")
	}

	// Create a couple of tenants
	app := NewApp()

	// Be sure to clean up the Synqly Accounts we create before exiting
	defer app.cleanup()

	if err := app.NewTenant(ctx, "Tenant ABC"); err != nil {
		log.Printf("unable to create tenant %v: %v", "Tenant ABC", err)
	}
	if err := app.NewTenant(ctx, "Tenant XYZ"); err != nil {
		log.Printf("unable to create tenant %v: %v", "Tenant XYZ", err)
	}

	// Configure one tenant to use event logging
	if err := app.configureEventLogging(ctx, "Tenant ABC"); err != nil {
		consoleLogger.Printf("ERROR: unable to configure event logging: %v", err)
	} else {
		go app.backgroundJob()

		// Wait for user to control c to exit
		select {}
	}

}
