package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"

	"github.com/synqly/go-sdk/client/engine"
	engineClient "github.com/synqly/go-sdk/client/engine/client"
	mgmt "github.com/synqly/go-sdk/client/management"
	mgmtClient "github.com/synqly/go-sdk/client/management/client"
)

var consoleLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

type App struct {
	Tenant *Tenant
	Config *Config
}

// NewApp instantiates a new App
func NewApp(config *Config) *App {
	return &App{
		Config: config,
	}
}

// A Tenant object represents one of your customers, as well as the state you
// would maintain for them in your application.
type Tenant struct {
	// ID: A unique identifier for one of your customers.
	ID string
	// SynqlyClient: A cached management client, used to manage Integrations.
	SynqlyClient *mgmtClient.Client
	// AssetMgmtClient: A cached engine client, used to interact with the asset management system.
	AssetMgmtClient *engineClient.Client
}

// NewTenant is used in a multi-tenant backend to create tenants. This would
// normally contain the logic you need to add a customer. When using Synqly,
// part of this logic should include creating an Account for them, and that is
// what's done here.
//
// Returns an error if a tenant with the same ID already exists or if a Synqly
// Account cannot be created for the Tenant.
func (a *App) NewTenant(ctx context.Context, id string) error {
	matched, err := regexp.MatchString(`^[a-z_-]*[a-z][a-z_-]*$`, id)
	if err != nil {
		return fmt.Errorf("error validating tenant ID: %w", err)
	}
	if !matched {
		return fmt.Errorf("invalid tenant ID: %v. Only letters, hyphens, and underscores are allowed", id)
	}

	// Do not allow duplicate tenant names
	if a.Tenant != nil && a.Tenant.ID == id {
		return fmt.Errorf("duplicate tenant name %v", id)
	}

	// Create a Synqly Client that can be used to interact with the tenant
	// We will use this client to create an Account and set up an Integration with a Asset Management system.
	client := mgmtClient.NewClient(
		mgmtClient.WithToken(a.Config.SynqlyOrgToken),
	)

	// Create a Synqly Account for this tenant
	_, err = client.Accounts.Create(ctx, &mgmt.CreateAccountRequest{
		Name: &id,
	})
	if err != nil {
		return fmt.Errorf("unable to create account: %w", err)
	}

	// Store the Tenant and associated Synqly objects in an in-memory cache.
	a.Tenant = &Tenant{
		ID:              id,
		SynqlyClient:    client,
		AssetMgmtClient: nil,
	}
	return nil
}

func (a *App) inmemConfig() *mgmt.CreateIntegrationRequest {
	return &mgmt.CreateIntegrationRequest{
		Fullname:       engine.String("In-Memory Connector"),
		ProviderConfig: &mgmt.ProviderConfig{SiemMockSiem: &mgmt.SiemMock{}},
	}
}

func (app *App) cleanup() {
	// Clean up the Accounts created in Synqly. In your application, you would
	// persist the Synqly Account id and Integration tokens to handle process
	// restarts. We are not doing that in this example, so we need to clean up
	// the accounts we created in Synqly.
	consoleLogger.Println("Cleaning up Synqly Account")
	ctx := context.Background()

	// Deleting the account will delete all credentials and integrations associated with the account.
	if err := app.Tenant.SynqlyClient.Accounts.Delete(ctx, app.Tenant.ID); err != nil {
		consoleLogger.Printf("Error deleting account %s: %s\n", app.Tenant.ID, err)
	}

	os.Exit(0)
}

func (a *App) configureAssetManagement(ctx context.Context, assetProvider string) error {
	var integrationReq *mgmt.CreateIntegrationRequest
	switch assetProvider {
	case "service-now":
		integrationReq = &mgmt.CreateIntegrationRequest{
			Fullname: engine.String("Asset Management Connector"),
			ProviderConfig: &mgmt.ProviderConfig{
				AssetsServicenow: &mgmt.AssetsServiceNow{
					Url: a.Config.ServiceNowURL,
					Credential: &mgmt.ServiceNowCredential{
						Token: &mgmt.TokenCredential{
							Secret: a.Config.ServiceNowTokenID,
						},
					},
				},
			},
		}
	case "armis":
		integrationReq = &mgmt.CreateIntegrationRequest{
			Fullname: engine.String("Asset Management Connector"),
			ProviderConfig: &mgmt.ProviderConfig{
				AssetsArmisCentrix: &mgmt.AssetsArmisCentrix{
					Url: a.Config.ArmisURL,
					Credential: &mgmt.ArmisCredential{
						TokenId: a.Config.ArmisTokenID,
					},
				},
			},
		}
	case "inmem":
		integrationReq = a.inmemConfig()
	default:
		return fmt.Errorf("invalid asset management provider: %v", assetProvider)
	}

	// Create an Integration in Synqly for the Asset Management system
	integration, err := a.Tenant.SynqlyClient.Integrations.Create(ctx, a.Tenant.ID, integrationReq)
	if err != nil {
		return fmt.Errorf("unable to create integration: %w", err)
	}

	// Store the Asset Management client in the Tenant cache
	a.Tenant.AssetMgmtClient = engineClient.NewClient(
		engineClient.WithToken(integration.Result.Token.Access.Secret),
	)

	return nil
}

func getDevices(ctx context.Context, client *engineClient.Client) ([]engine.Device, error) {
	req := &engine.QueryDevicesRequest{
		Limit: engine.Int(1),
		Order: engine.String("device.name[desc]"),
	}

	queryDevicesResponse, err := client.Assets.QueryDevices(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to query devices: %w", err)
	}

	consoleLogger.Println("Devices queried successfully")

	return queryDevicesResponse.Result, nil
}

func createDevice(ctx context.Context, client *engineClient.Client, device engine.Device) error {
	deviceCopy := *device
	newName := *device.Device.Name + "-copy"
	newUid := *device.Device.Uid + "-copy"
	deviceCopy.Device.Name = &newName
	deviceCopy.Device.Uid = &newUid
	req := &engine.CreateDeviceRequest{
		Device: &deviceCopy,
	}

	_, err := client.Assets.CreateAsset(ctx, req)
	if err != nil {
		return fmt.Errorf("unable to create device: %w", err)
	}

	consoleLogger.Println("")
	consoleLogger.Println("Device created successfully")
	return nil
}

func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	ctx := context.Background()

	// Load config variables from the env file
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	if config.ServiceNowURL == "" || config.ServiceNowTokenID == "" {
		log.Fatal("service-now URL and token ID are required")
	}

	app := NewApp(&config)

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

	customerName := "asset_demo_customer"
	consoleLogger.Printf("Creating %s tenant\n", customerName)
	if err := app.NewTenant(ctx, customerName); err != nil {
		log.Fatal(err)
	}

	err = app.configureAssetManagement(ctx, "service-now")
	if err != nil {
		log.Fatal(err)
	}

	// Query devices
	devices, err := getDevices(ctx, app.Tenant.AssetMgmtClient)
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range devices {
		consoleLogger.Printf("Device: %v\n", device)
	}

	// Create a device
	if len(devices) > 0 {
		err = createDevice(ctx, app.Tenant.AssetMgmtClient, devices[0])
		if err != nil {
			log.Fatal(err)
		}
	}
}
