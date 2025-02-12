package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/synqly/go-sdk/client/engine"
	"github.com/synqly/go-sdk/examples/common"
	"log"
	"os"

	mgmt "github.com/synqly/go-sdk/client/management"
)

var consoleLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	// Load config variables from the env file, and if that is not present, then try env vars
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	if config.synqlyOrgToken == "" || config.edrToken == "" || config.edrURL == "" {
		log.Fatal("Must set following environment variables: SYNQLY_ORG_TOKEN EDR_TOKEN EDR_URL")
	}

	if err := demoActions(config); err != nil {
		log.Fatal(err)
	}
}
func ProviderConfig(c Config) *mgmt.ProviderConfig {
	return &mgmt.ProviderConfig{
		EdrSentinelone: &mgmt.EdrSentinelOne{
			Url: c.edrURL,
			Credential: &mgmt.SentinelOneCredential{
				Token: &mgmt.TokenCredential{Secret: c.edrToken},
			},
		},
	}
}

func demoActions(config Config) error {
	ctx := context.Background()

	id := "EDR SentinelOne Tenant"
	t, err := common.NewTenant(ctx, id, "tenant_store.yaml", config.synqlyOrgToken, map[mgmt.CategoryId]*mgmt.CreateIntegrationRequest{
		mgmt.CategoryIdEdr: {
			Fullname:       mgmt.String("SentinelOne Identity Provider"),
			ProviderConfig: ProviderConfig(config),
		},
	})
	if err != nil {
		return fmt.Errorf("unable to create tenant: %w", err)
	}

	consoleLogger.Print("Tenant loaded successfully")

	// Get some applications
	var resp *engine.QueryApplicationsResponse
	page := 0
	for {
		page++
		if resp != nil && resp.Cursor != "" {
			resp, err = t.Synqly.EngineClients[mgmt.CategoryIdEdr].Edr.QueryApplications(ctx, &engine.QueryApplicationsRequest{Cursor: &resp.Cursor})
		} else {
			resp, err = t.Synqly.EngineClients[mgmt.CategoryIdEdr].Edr.QueryApplications(ctx, &engine.QueryApplicationsRequest{Limit: engine.Int(2)})

		}
		if err != nil {
			return fmt.Errorf("unable to query applications page %d: %w", page, err)
		}

		for _, app := range resp.Result {
			apps, _ := json.MarshalIndent(app, "", "  ")
			fmt.Println()
			consoleLogger.Printf("apps %s:\n", apps)
		}

		if resp.Cursor == "" || page == 5 {
			// no more data to fetch
			break
		}
	}

	// Get an endpoint
	endpointRes, err := t.Synqly.EngineClients[mgmt.CategoryIdEdr].Edr.QueryEndpoints(ctx, &engine.QueryEndpointsRequest{})
	if err != nil {
		return fmt.Errorf("unable to query endpoints: %w", err)
	}
	for _, app := range endpointRes.Result {
		apps, _ := json.MarshalIndent(app, "", "  ")
		fmt.Println()
		consoleLogger.Printf("agents %s:\n", apps)
	}

	// Quarantine an endpoint, uncomment to execute

	// consoleLogger.Printf("Quarantining endpoint %s", *endpointRes.Result[0].Device.Uid)
	// endpointId := *endpointRes.Result[0].Device.Uid
	// t.Synqly.EngineClients[mgmt.CategoryIdEdr].Edr.NetworkQuarantine(ctx, &engine.NetworkQuarantineRequest{
	// 	EndpointIds: []string{endpointId},
	// 	State:       engine.ConnectionStateDisconnect,
	// })

	// t.Synqly.EngineClients[mgmt.CategoryIdEdr].Edr.NetworkQuarantine(ctx, &engine.NetworkQuarantineRequest{
	// 	EndpointIds: []string{endpointId},
	// 	State:       engine.ConnectionStateConnect,
	// })

	return nil
}
