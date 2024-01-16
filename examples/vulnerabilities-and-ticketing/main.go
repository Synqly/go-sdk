package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	koanfyaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"

	"github.com/synqly/go-sdk/client/engine"
	engineClient "github.com/synqly/go-sdk/client/engine/client"
	"github.com/synqly/go-sdk/client/engine/ocsf/securityfinding"
	mgmt "github.com/synqly/go-sdk/client/management"
	"github.com/synqly/go-sdk/examples/common"
)

var (
	k = koanf.New(".")
)

var consoleLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

// notification creates a ticket in the tenant's ticketing connector for the given vulnerability finding
func notification(ctx context.Context, client *engineClient.Client, finding *securityfinding.SecurityFinding) error {
	tags := []string{fmt.Sprintf("asset:%s", *finding.Resources[0].Uid)}

	// search for existing ticket
	tickets, err := client.Ticketing.QueryTickets(ctx, &engine.QueryTicketsRequest{
		Filter: []*string{
			engine.String(fmt.Sprintf("tags[in]%s", strings.Join(tags, ","))),
		},
	})
	if err != nil {
		return err
	}
	for _, ticket := range tickets.Result {
		consoleLogger.Printf("Found existing ticket with id %s\n", ticket.Id)
		return nil
	}

	title := fmt.Sprintf("Vulnerability found for %s", *finding.Resources[0].Name)
	summary := fmt.Sprintf("%s has a %s priority vulnerability", *finding.Resources[0].Name, strings.ToLower(*finding.Severity))
	descrTemplate := `%s found a vulnerability: %s
%s

%s

h3. Asset
%s

h3. Remediation
%s

h3. CVE information
CVSS base score: %f

See also:
%s
`

	asset, err := json.MarshalIndent(finding.Resources[0].Data, "", "  ")
	if err != nil {
		return err
	}

	description := fmt.Sprintf(descrTemplate,
		finding.Metadata.Product.VendorName,
		finding.Finding.Title,
		*finding.Finding.Desc,
		*finding.Message,
		asset,
		*finding.Finding.Remediation.Desc,
		finding.Vulnerabilities[0].Cve.Cvss.BaseScore,
		strings.Join(finding.Vulnerabilities[0].References, "\n"),
	)

	priority := engine.PriorityLow
	switch *finding.Severity {
	case "Critical":
		priority = engine.PriorityUrgent
	case "High":
		priority = engine.PriorityHigh
	case "Medium":
		priority = engine.PriorityMedium
	}

	tick, err := client.Ticketing.CreateTicket(ctx, &engine.CreateTicketRequest{
		Name:        title,
		Summary:     summary,
		Description: engine.String(description),
		IssueType:   engine.String("Bug"),
		Priority:    priority.Ptr(),
		Project:     engine.String("Test"),
		Tags:        tags,
	})
	if err != nil {
		return err
	}

	consoleLogger.Printf("Created a ticket with id %s\n", tick.Result.Id)

	return nil
}

func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	parser := koanfyaml.Parser()
	if err := k.Load(file.Provider("config.yaml"), parser); err != nil {
		k.Load(env.Provider("SYNQLY_", "_", func(s string) string {
			return strings.ToLower(s)
		}), nil)
	}

	synqlyOrgToken := k.String("synqly.org.token")
	tenableToken := k.String("synqly.tenable.token")
	jiraUser := k.String("synqly.jira.user")
	jiraToken := k.String("synqly.jira.token")

	if synqlyOrgToken == "" || tenableToken == "" || jiraUser == "" || jiraToken == "" {
		log.Fatal("Missing required config")
	}

	ctx := context.Background()

	t, err := common.NewTenant(ctx, "Acme Corp", "tenant_store.yaml", synqlyOrgToken, map[mgmt.ProviderId]common.AuthAndConfig{
		"ticketing": {
			Auth: &mgmt.CreateCredentialRequest{
				Name: engine.String("jira"),
				Config: mgmt.NewCredentialConfigFromBasic(&mgmt.BasicCredential{
					Username: jiraUser,
					Secret:   jiraToken,
				}),
			},
			Config: &mgmt.CreateIntegrationRequest{
				Fullname:     engine.String("Ticketing Connector"),
				Category:     "ticketing",
				ProviderType: "jira",
				ProviderConfig: mgmt.NewProviderConfigFromTicketing(&mgmt.TicketingConfig{
					Endpoint: engine.String("https://synqly.atlassian.net"),
				}),
			},
		},
		"vulnerabilities": {
			Auth: &mgmt.CreateCredentialRequest{
				Name: engine.String("tenable"),
				Config: mgmt.NewCredentialConfigFromToken(&mgmt.TokenCredential{
					Secret: tenableToken,
				}),
			},
			Config: &mgmt.CreateIntegrationRequest{
				Fullname:     engine.String("Vulnerability Scanner"),
				Category:     "vulnerabilities",
				ProviderType: "tenable",
				ProviderConfig: mgmt.NewProviderConfigFromVulnerabilities(&mgmt.VulnerabilityConfig{
					Endpoint: engine.String("https://cloud.tenable.com"),
				}),
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	findings, err := t.Synqly.EngineClients["vulnerabilities"].Vulnerabilities.QueryVulnerabilityFindings(ctx, &engine.QueryFindingsRequest{
		Filter: []*string{
			engine.String("severity[in]Critical,High,Medium"),
			engine.String("finding.last_seen_time[gte]2024-01-05T00:00:00Z"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	consoleLogger.Printf("Found %d security findings from Tenable\n", len(findings.Result))

	for _, finding := range findings.Result {
		if err := notification(ctx, t.Synqly.EngineClients["ticketing"], finding.SecurityFinding); err != nil {
			log.Fatal(err)
		}
	}
}
