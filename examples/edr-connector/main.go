package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/knadh/koanf"
	koanfyaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/synqly/go-sdk/client/engine"
	mgmt "github.com/synqly/go-sdk/client/management"
	"github.com/synqly/go-sdk/examples/common"
)

var (
	k               = koanf.New(".")
	sentinelOneConf = sentinelOneConfig{}
)

var consoleLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

type sentinelOneConfig struct {
	URL   string `koanf:"url"`
	Token string `koanf:"token"`
}

func (c *sentinelOneConfig) Validate() error {
	if c.URL == "" {
		return errors.New("sentinelone.url is required")
	}
	if c.Token == "" {
		return errors.New("sentinelone.token is required")
	}
	return nil
}

type sentinelOneProvider struct {
	config *sentinelOneConfig
}

func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	// load config -- try from a config file, and if that is not present, then try env vars
	parser := koanfyaml.Parser()
	if err := k.Load(file.Provider("config-edr.yaml"), parser); err != nil {
		err := k.Load(env.Provider("SYNQLY_", "_", func(s string) string {
			return strings.ToLower(s)
		}), nil)
		if err != nil {
			log.Fatalf("unable to load config-edr.yaml or env vars: %v", err)
		}
	}

	orgToken := k.String("synqly.org.token")
	if orgToken == "" {
		log.Fatal("synqly.org.token is required. Set SYNQLY_ORG_TOKEN or add it to config-edr.yaml.")
	}

	err := k.Unmarshal("synqly.sentinelone", &sentinelOneConf)
	if err != nil {
		log.Fatal(err)
	}

	if err := sentinelOneConf.Validate(); err != nil {
		log.Fatalf("invalid sentinelone config: %s. Set SentinelOne configuration through env vars or add it to config-edr.yaml", err)
	}

	sentinelOneProvider := &sentinelOneProvider{config: &sentinelOneConf}

	if err := sentinelOneProvider.demoActions(orgToken, &sentinelOneConf); err != nil {
		log.Fatal(err)
	}
}

func (s *sentinelOneProvider) demoActions(orgToken string, sentinelOneConf *sentinelOneConfig) error {
	ctx := context.Background()

	id := "EDR SentinelOne Tenant"
	t, err := common.NewTenant(ctx, id, "tenant_store.yaml", orgToken, map[mgmt.CategoryId]*mgmt.CreateIntegrationRequest{
		mgmt.CategoryIdEdr: {
			Fullname: mgmt.String("SentinelOne Identity Provider"),
			ProviderConfig: &mgmt.ProviderConfig{
				EdrSentinelone: &mgmt.EdrSentinelOne{
					Url: sentinelOneConf.URL,
					Credential: &mgmt.SentinelOneCredential{
						Token: &mgmt.TokenCredential{Secret: sentinelOneConf.Token},
					},
				},
			},
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
			consoleLogger.Println()
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
		consoleLogger.Println()
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
