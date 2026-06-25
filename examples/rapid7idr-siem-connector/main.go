package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/knadh/koanf"
	"github.com/synqly/go-sdk/examples/common"

	"github.com/synqly/go-sdk/client/engine"
	mgmt "github.com/synqly/go-sdk/client/management"
)

var (
	k                 = koanf.New(".")
	rapid7IdrSiemConf = rapid7IdrSiemConfig{}
)

var consoleLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

type rapid7IdrSiemConfig struct {
	URL   string `koanf:"url"`
	Token string `koanf:"token"`
}

func (c *rapid7IdrSiemConfig) Validate() error {
	if c.URL == "" {
		return errors.New("rapid7idr.url is required")
	}
	if c.Token == "" {
		return errors.New("rapid7idr.token is required")
	}
	return nil
}

type rapid7IdrProvider struct {
	config *rapid7IdrSiemConfig
}

func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	orgToken := "SYNQLY_ORG_TOKEN"
	if orgToken == "" {
		log.Fatal("synqly.token is required. Set SYNQLY_ORG_TOKEN or add it to config.yaml.")
	}

	rapid7IdrSiemConf.Token = "RAPID7_IDR_TOKEN"
	rapid7IdrSiemConf.URL = "RAPID7_IDR_URL"

	if err := rapid7IdrSiemConf.Validate(); err != nil {
		log.Fatalf("invalid rapid7idr config: %s. Set Rapid7Idr configuration through env vars or add it to config.yaml", err)
	}

	rapid7IdrProvider := &rapid7IdrProvider{config: &rapid7IdrSiemConf}

	if err := rapid7IdrProvider.demoActions(orgToken, &rapid7IdrSiemConf); err != nil {
		log.Fatal(err)
	}
}

func (s *rapid7IdrProvider) demoActions(orgToken string, rapid7IdrSiemConf *rapid7IdrSiemConfig) error {
	ctx := context.Background()

	id := "Rapid7 Idr SIEM Tenant"
	t, err := common.NewTenant(ctx, id, "tenant_store.yaml", orgToken, map[mgmt.CategoryId]*mgmt.CreateIntegrationRequest{
		mgmt.CategoryIdSiem: {
			Fullname: mgmt.String("Rapi7IdrSiem Identity Provider "),
			ProviderConfig: mgmt.NewProviderConfigFromSiemRapid7Insightidr(&mgmt.SiemRapid7InsightIdr{
				Url: rapid7IdrSiemConf.URL,
				Credential: mgmt.NewRapid7InsightCloudCredentialFromToken(&mgmt.TokenCredential{
					Secret: rapid7IdrSiemConf.Token,
				}),
			}),
		},
	})
	if err != nil {
		return fmt.Errorf("unable to create tenant: %w", err)
	}

	// Get Investigations, including raw data
	investigationId := "rrn:investigation:us2:82ca50e6-3684-4814-bb61-c8ade4e2b94e:investigation:PNT21UETP6EJ"
	resp, err := t.Synqly.EngineClients[mgmt.CategoryIdSiem].Siem.GetInvestigation(ctx, investigationId, &engine.GetInvestigationRequest{
		IncludeRawData: engine.Bool(true),
	})

	if err != nil {
		return err
	}

	fmt.Print(resp.Result.Id)
	fmt.Print(resp.Result.RawData)

	// Get Evidence for an Investigation, including raw data
	investigationIdForEvidence := "rrn:investigation:us2:82ca50e6-3684-4814-bb61-c8ade4e2b94e:investigation:G9HB17N4PSCV"

	evidenceResp, err := t.Synqly.EngineClients[mgmt.CategoryIdSiem].Siem.GetEvidence(ctx, investigationIdForEvidence, &engine.GetInvestigationEvidenceRequest{
		IncludeRawData: engine.Bool(true),
	})

	if err != nil {
		return err
	}
	fmt.Print(evidenceResp.Result.RawData)

	// Patch Investigation - always returns RawData in the response

	patchData := []map[string]interface{}{
		{"op": "replace", "path": "/status", "value": "CLOSED"},
	}

	investigationIdToClose := "rrn:investigation:us2:82ca50e6-3684-4814-bb61-c8ade4e2b94e:investigation:PNT21UETP6EJ"
	patchRes, err := t.Synqly.EngineClients[mgmt.CategoryIdSiem].Siem.PatchInvestigation(ctx, investigationIdToClose, patchData)
	if err != nil {
		return err
	}

	fmt.Print(patchRes.Result.RawData)

	// QueryEvents (logs), for Rapid7 LEQL, use PassthroughParam & the flag for including raw data so we get raw results in our response
	// Query logset_name "Audit Logs" for logs containing "postman" in the last 14 days
	queryEventsReq := &engine.QuerySiemEventsRequest{
		PassthroughParam: []*string{
			engine.String(`query:where("postman")`),
			engine.String("time_range:last 1 hour"),
			engine.String("logset_name:Audit Logs"),
		},
		IncludeRawData: engine.Bool(true),
	}

	queryEventsRes, err := t.Synqly.EngineClients[mgmt.CategoryIdSiem].Siem.QueryEvents(ctx, queryEventsReq)

	if err != nil {
		return err
	}

	for _, event := range queryEventsRes.Result {
		fmt.Print(event["raw_data"])
	}

	return nil
}
