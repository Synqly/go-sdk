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

func ticketingProviderConfig(jiraUser, jiraToken string) (*mgmt.CreateIntegrationRequest, error) {
	if jiraUser == "" || jiraToken == "" {
		return nil, fmt.Errorf("missing ticketing provider config")
	}

	integrationReq := &mgmt.CreateIntegrationRequest{
		Fullname: engine.String("Ticketing Connector"),
		ProviderConfig: mgmt.NewProviderConfigFromTicketingJira(&mgmt.TicketingJira{
			Url: "https://synqly.atlassian.net",
			Credential: mgmt.NewJiraCredentialFromBasic(&mgmt.BasicCredential{
				Username: jiraUser,
				Secret:   jiraToken,
			}),
		}),
	}

	return integrationReq, nil
}

func vulnerabilityProviderConfig(tenableToken, qualysEndpoint, qualysUser, qualysSecret string) (*mgmt.CreateIntegrationRequest, error) {
	if tenableToken != "" {
		return &mgmt.CreateIntegrationRequest{
			Fullname: engine.String("Vulnerability Scanner"),
			ProviderConfig: mgmt.NewProviderConfigFromVulnerabilitiesTenableCloud(&mgmt.VulnerabilitiesTenableCloud{
				Url: mgmt.String("https://cloud.tenable.com"),
				Credential: mgmt.NewTenableCloudCredentialFromToken(&mgmt.TokenCredential{
					Secret: tenableToken,
				}),
			}),
		}, nil
	}

	if qualysEndpoint != "" && qualysUser != "" && qualysSecret != "" {
		return &mgmt.CreateIntegrationRequest{
			Fullname: engine.String("Vulnerability Scanner"),
			ProviderConfig: mgmt.NewProviderConfigFromVulnerabilitiesQualysCloud(&mgmt.VulnerabilitiesQualysCloud{
				Url: qualysEndpoint,
				Credential: mgmt.NewQualysCloudCredentialFromBasic(&mgmt.BasicCredential{
					Username: qualysUser,
					Secret:   qualysSecret,
				}),
			}),
		}, nil
	}

	return nil, fmt.Errorf("missing vulnerability provider config")
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
	jiraUser := k.String("synqly.jira.user")
	jiraToken := k.String("synqly.jira.token")
	tenableToken := k.String("synqly.tenable.token")
	qualysEndpoint := k.String("synqly.qualys.endpoint")
	qualysUser := k.String("synqly.qualys.user")
	qualysSecret := k.String("synqly.qualys.secret")

	if synqlyOrgToken == "" {
		log.Fatal("Missing Synqly Org token")
	}

	ticketingProvider, err := ticketingProviderConfig(jiraUser, jiraToken)
	if err != nil {
		log.Fatal(err)
	}

	vulnProvider, err := vulnerabilityProviderConfig(tenableToken, qualysEndpoint, qualysUser, qualysSecret)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	t, err := common.NewTenant(ctx, "Acme Corp", "tenant_store.yaml", synqlyOrgToken, map[mgmt.CategoryId]*mgmt.CreateIntegrationRequest{
		mgmt.CategoryIdTicketing:       ticketingProvider,
		mgmt.CategoryIdVulnerabilities: vulnProvider,
	})
	if err != nil {
		log.Fatal(err)
	}

	consoleLogger.Printf("Using %s as vulnerability provider\n", vulnProvider.ProviderConfig.Type)

	findings, err := t.Synqly.EngineClients["vulnerabilities"].Vulnerabilities.QueryFindings(ctx, &engine.QueryFindingsRequest{
		Filter: []*string{
			engine.String("severity[in]Critical,High,Medium"),
			engine.String("finding.last_seen_time[gte]2024-01-05T00:00:00Z"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	consoleLogger.Printf("Found %d security findings from %s\n", len(findings.Result), vulnProvider.ProviderConfig.Type)

	for _, finding := range findings.Result {
		if err := notification(ctx, t.Synqly.EngineClients["ticketing"], finding.SecurityFinding); err != nil {
			log.Fatal(err)
		}
	}
}
