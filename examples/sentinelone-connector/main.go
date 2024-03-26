package main

import (
	"bufio"
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
	"github.com/synqly/go-sdk/examples/common"

	mgmt "github.com/synqly/go-sdk/client/management"
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
		return errors.New("sentielone.url is required")
	}
	if c.Token == "" {
		return errors.New("sentielone.token is required")
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
	if err := k.Load(file.Provider("config.yaml"), parser); err != nil {
		k.Load(env.Provider("SYNQLY_", "_", func(s string) string {
			return strings.ToLower(s)
		}), nil)
	}

	orgToken := k.String("synqly.token")
	if orgToken == "" {
		log.Fatal("synqly.token is required. Set SYNQLY_ORG_TOKEN or add it to config.yaml.")
	}

	err := k.Unmarshal("sentinelone", &sentinelOneConf)
	if err != nil {
		log.Fatal(err)
	}

	if err := sentinelOneConf.Validate(); err != nil {
		log.Fatalf("invalid sentinelone config: %s. Set SentinelOne configuration through env vars or add it to config.yaml", err)
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
			ProviderConfig: mgmt.NewProviderConfigFromEdrSentinelone(&mgmt.EdrSentinelOne{
				Url: sentinelOneConf.URL,
				Credential: mgmt.NewSentinelOneCredentialFromToken(&mgmt.TokenCredential{
					Secret: sentinelOneConf.Token,
				}),
			}),
		},
	})
	if err != nil {
		return fmt.Errorf("unable to create tenant: %w", err)
	}

	// Get some applications
	resp, err := t.Synqly.EngineClients[mgmt.CategoryIdEdr].Edr.QueryApplications(ctx, &engine.QueryApplicationsRequest{Limit: engine.Int(2)})

	if err != nil {
		return fmt.Errorf("error querying applications: %s", err)
	}
	for _, app := range resp.Result {
		apps, _ := json.MarshalIndent(app, "", "  ")
		fmt.Println()
		consoleLogger.Printf("apps %s:\n", apps)
	}

	// Get some more applications
	resp, err = t.Synqly.EngineClients[mgmt.CategoryIdEdr].Edr.QueryApplications(ctx, &engine.QueryApplicationsRequest{Cursor: &resp.Cursor})

	if err != nil {
		return fmt.Errorf("error querying applications: %s", err)
	}
	for _, app := range resp.Result {
		apps, _ := json.MarshalIndent(app, "", "  ")
		fmt.Println()
		consoleLogger.Printf("apps %s:\n", apps)
	}

	// Get an endpoint
	endpointRes, err := t.Synqly.EngineClients[mgmt.CategoryIdEdr].Edr.QueryEndpoints(ctx, &engine.QueryEndpointsRequest{})

	if err != nil {
		return err
	}
	for _, app := range endpointRes.Result {
		apps, _ := json.MarshalIndent(app, "", "  ")
		fmt.Println()
		consoleLogger.Printf("agents %s:\n", apps)
	}

	consoleLogger.Printf("Quarantining endpoint %s", *endpointRes.Result[0].Device.Uid)

	endpointId := *endpointRes.Result[0].Device.Uid
	t.Synqly.EngineClients[mgmt.CategoryIdEdr].Edr.NetworkQuarantine(ctx, &engine.NetworkQuarantineRequest{
		EndpointIds: []string{endpointId},
		State:       engine.ConnectionStateDisconnect,
	})

	fmt.Println("press any key to add endpoint back to network (Unquarantine)...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	t.Synqly.EngineClients[mgmt.CategoryIdEdr].Edr.NetworkQuarantine(ctx, &engine.NetworkQuarantineRequest{
		EndpointIds: []string{endpointId},
		State:       engine.ConnectionStateConnect,
	})

	return nil
}
