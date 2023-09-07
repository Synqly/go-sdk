package main

import (
	"context"
	"fmt"
	"time"

	"github.com/synqly/go-sdk/client/engine"
	"github.com/synqly/go-sdk/client/engine/events"
	"github.com/synqly/go-sdk/client/management"
)

const (
	synqlyOrgId    = "synqly-org-id"
	synqlyOrgToken = "synqly-org-token"
)

type App struct {
	Tenants []*Tenant
}

func NewApp() *App {
	return &App{}
}

type Tenant struct {
	ID              string
	SynqlyAccountId string
	SynqlyClient    management.Client
	EventLogger     engine.Client
}

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

	// Typically this info would be collected from the user in the UI and then
	// the UI would call the Synqly API to create the Integration. We are doing
	// it here to demonstrate how it works.
	splunkURL := "https://my-org.splunkcloud.com:8088/services/collector/event"
	splunkToken := "my-splunk-token"

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

func backgroundJob(app *App) {
	// Do some work for each tenant every 10 seconds
	for {
		for _, tenant := range app.Tenants {
			fmt.Printf("Doing some important work for tenant %s\n", tenant.ID)

			// Log the result of the work to Synqly
			if tenant.EventLogger != nil {
				err := tenant.EventLogger.Events().PostEvent(context.Background(),
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
								VendorName: "Synqly",
							},
							Version: "1.0.0",
						},
						Time:       int(time.Now().UTC().Unix()),
						SeverityId: events.SeverityIdInformational,
						TypeUid:    events.ScheduledJobActivityTypeUidScheduledJobActivityUpdate,
					}))
				if err != nil {
					fmt.Printf("Error logging event for tenant %s: %s\n", tenant.ID, err)
				}
			}
		}

		time.Sleep(10 * time.Second)
	}

}

func main() {
	ctx := context.Background()

	// Create a couple of tenants
	app := NewApp()
	if err := app.NewTenant(ctx, "Tenant ABC"); err != nil {
		panic(err)
	}
	if err := app.NewTenant(ctx, "Tenant XYZ"); err != nil {
		panic(err)
	}

	// Configure one tenant to use event logging
	if err := app.configureEventLogging(ctx, "Tenant ABC"); err != nil {
		panic(err)
	}

	go backgroundJob(app)

	// wait for user to control c to exit
	select {}
}
