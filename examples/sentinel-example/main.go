package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"

	engine "github.com/synqly/go-sdk/client/engine"
	engineClient "github.com/synqly/go-sdk/client/engine/client"
	mgmt "github.com/synqly/go-sdk/client/management"
	mgmtClient "github.com/synqly/go-sdk/client/management/client"
)

func main() {
	ctx := context.Background()

	// Load config variables from the env file
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	if config.synqlyOrgToken == "" || config.azureClientId == "" || config.azureClientSecret == "" {
		log.Fatal("Must set following environment variables: SYNQLY_ORG_TOKEN AZURE_CLIENT_ID AZURE_CLIENT_SECRET")
	}

	client := mgmtClient.NewClient(
		mgmtClient.WithToken(config.synqlyOrgToken),
	)

	integrationName := "alerts"

	defer func() {
		if err := client.Integrations.Delete(ctx, "acme-corp", integrationName); err != nil {
			log.Printf("Failed to delete integration: %v", err)
		} else {
			log.Print("Integration deleted")
		}
	}()

	integration, err := client.Integrations.Create(ctx, "acme-corp", &mgmt.CreateIntegrationRequest{
		Name: mgmt.String(integrationName),
		ProviderConfig: &mgmt.ProviderConfig{
			SinkAzureMonitorLogs: &mgmt.SinkAzureMonitorLogs{
				Url: config.dceURL,
				Credential: &mgmt.AzureMonitorLogsCredential{
					Token: &mgmt.TokenCredential{
						Secret: config.azureClientSecret,
					},
				},
				ClientId:   config.azureClientId,
				TenantId:   config.azureTenantId,
				RuleId:     config.dcrId,
				StreamName: "Custom-CommonSecurityLog",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Integration created: %s", integration.Result.Integration.Name)

	e := engineClient.NewClient(
		engineClient.WithToken(integration.Result.Token.Access.Secret),
	)

	detection, err := sampleDetectionFinding()
	if err != nil {
		log.Fatal(err)
	}

	err = e.Sink.PostEvents(ctx, []*engine.Event{
		detection,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Detection finding sent")
}

func sampleDetectionFinding() (*engine.Event, error) {
	s := fmt.Sprintf(`{
        "class_name": "Detection Finding",
        "activity_id": 1,
        "category_name": "Findings",
        "category_uid": 2,
        "class_uid": 2004,
        "count": 1,
        "message": "%s",
        "metadata": {
            "log_provider": "Uploaded",
            "product": {
                "name": "%s",
                "vendor_name": "Synqly"
            },
            "version": "1.1.0"
        },
        "severity_id": 4,
        "severity": "High",
        "status_detail": "%s",
        "status_id": 1,
        "start_time": %d,
        "time": %d,
        "type_uid": 200401,
        "finding_info": {
            "related_analytics": [
                {
                    "category": "analytic_category",
                    "desc": "analytic_description",
                    "name": "analytic_name",
                    "type": "Rule",
                    "type_id": 1,
                    "uid": "analytic_uid",
                    "version": "analytic_version"
                }
            ],
            "created_time": %d,
            "desc": "Description of the detection finding",
            "attacks": [
                {
                    "sub_technique": {
                        "uid": "T1234.001",
                        "name": "Sub Technique Name",
                        "src_url": "https://attack.mitre.org/techniques/T1234/001/"
                    },
                    "tactic": {
                        "uid": "TA0001",
                        "name": "Tactic Name",
                        "src_url": "https://attack.mitre.org/tactics/TA0001/"
                    },
                    "technique": {
                        "uid": "T1234",
                        "name": "Technique Name",
                        "src_url": "https://attack.mitre.org/techniques/T1234/"
                    },
                    "version": "7.0"
                }
            ]
        },
        "observables": [
            {
                "name": "observable_name",
                "type": "Fingerprint",
                "type_id": 30,
                "value": "observable_value",
                "reputation": {
                    "provider": "Reputation Provider",
                    "base_score": 7.1,
                    "score": "ProbablyMalicious",
                    "score_id": 9
                }
            }
        ]
    }`, gofakeit.HackerPhrase(), gofakeit.Company(), gofakeit.BeerName(), time.Now().UTC().Unix(), time.Now().UTC().Unix(), time.Now().UTC().Unix())

	event := engine.Event{}
	if err := json.Unmarshal([]byte(s), &event); err != nil {
		return nil, err
	}

	return &event, nil
}
