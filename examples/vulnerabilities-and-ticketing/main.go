package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/synqly/go-sdk/client/engine"
	engineClient "github.com/synqly/go-sdk/client/engine/client"
	"github.com/synqly/go-sdk/client/engine/ocsf/securityfinding"
	mgmt "github.com/synqly/go-sdk/client/management"
	mgmtClient "github.com/synqly/go-sdk/client/management/client"
)

var (
	synqlyOrgToken = os.Getenv("SYNQLY_ORG_TOKEN")

	// Tenable is used as an example for vulnerability scanner. For each
	// vulnerability finding, a ticket will be created in jira
	tenableToken = os.Getenv("TENABLE_TOKEN")
	jiraToken    = os.Getenv("JIRA_TOKEN") // Jira is used as an example ticketing system
)

var consoleLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

// Tenant represents a customer of yours. It has its own account in Synqly, and this struct
// keeps the information you need to interact with Synqly on behalf of the customer.
type Tenant struct {
	ID     string
	synqly SynqlyRuntime
}

// SynqlyRuntime keeps the clients needed to use the Synqly management, vulnerability and ticketing APIs
type SynqlyRuntime struct {
	accountId       string
	management      *mgmtClient.Client
	vulnerabilities *engineClient.Client
	ticketing       *engineClient.Client
}

// TenantConfig is the configuration you need to save for your tenants to re-connect to Synqly
type TenantConfig struct {
	SynqlyConfig SynqlyConfig `yaml:"synqly_config"`
}

type SynqlyConfig struct {
	AccountID            string `yaml:"account_id"`
	VulnerabilitiesToken string `yaml:"vulnerabilities_token"`
	TicketingToken       string `yaml:"ticketing_token"`
}

// NewTenant creates a tenant and configures it with Synqly. If this tenant does not
// have a saved config file, then a new Synqly Account will be created with a vulnerability
// and ticketing connector. If the tenant already has a config file, then the existing
// Synqly Account and connectors will be used.
func NewTenant(ctx context.Context, id string, configFilePath string) (*Tenant, error) {
	tenant := &Tenant{}
	config := &TenantConfig{}

	configFile, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	if err := yaml.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, err
	}

	tenant.synqly.management = mgmtClient.NewClient(
		mgmtClient.WithAuthToken(synqlyOrgToken),
	)

	if err := tenant.Init(ctx, id, config); err != nil {
		return nil, err
	}

	// write config back to file
	configFile, err = os.Create(configFilePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	if err := yaml.NewEncoder(configFile).Encode(config); err != nil {
		return nil, err
	}

	return tenant, nil
}

// Init initializes the tenant's Synqly Account and Connectors clients
func (t *Tenant) Init(ctx context.Context, tenantID string, config *TenantConfig) error {
	accountId := config.SynqlyConfig.AccountID
	if accountId == "" {
		account, err := t.synqly.management.Accounts.CreateAccount(ctx, &mgmt.CreateAccountRequest{
			Name: "Synqly Vulnerabilities and Ticketing Demo",
		})
		if err != nil {
			return fmt.Errorf("init failed, unable to create account: %w", err)
		}
		accountId = account.Result.Account.Id
		consoleLogger.Printf("Created new account: %s", accountId)
	}
	config.SynqlyConfig.AccountID = accountId
	t.synqly.accountId = accountId

	consoleLogger.Printf("Using account with id: %s\n", accountId)

	var err error
	var token string
	if config.SynqlyConfig.VulnerabilitiesToken == "" {
		t.synqly.vulnerabilities, token, err = t.configureVulnerabilityConnector(ctx)
		if err != nil {
			return fmt.Errorf("init failed, unable to configure vulnerability connector: %w", err)
		}
		config.SynqlyConfig.VulnerabilitiesToken = token
	} else {
		t.synqly.vulnerabilities = engineClient.NewClient(
			engineClient.WithAuthToken(config.SynqlyConfig.VulnerabilitiesToken),
		)
	}

	if config.SynqlyConfig.TicketingToken == "" {
		t.synqly.ticketing, token, err = t.configureTicketingConnector(ctx)
		if err != nil {
			return fmt.Errorf("init failed, unable to configure ticketing connector: %w", err)
		}
		config.SynqlyConfig.TicketingToken = token
	} else {
		t.synqly.ticketing = engineClient.NewClient(
			engineClient.WithAuthToken(config.SynqlyConfig.TicketingToken),
		)
	}

	return nil
}

func (t *Tenant) createVulnerabilityConnector(ctx context.Context, providerType string, providerConfig *mgmt.VulnerabilityConfig) (*mgmt.CreateIntegrationResponseResult, error) {
	response, err := t.synqly.management.Integrations.CreateIntegration(ctx, t.synqly.accountId, &mgmt.CreateIntegrationRequest{
		Name:           "Vulnerability Connector",
		Category:       "vulnerabilities",
		ProviderType:   providerType,
		ProviderConfig: mgmt.NewProviderConfigFromVulnerabilities(providerConfig),
	})
	if err != nil {
		return nil, err
	}

	return response.Result, nil
}

func (t *Tenant) configureVulnerabilityConnector(ctx context.Context) (*engineClient.Client, string, error) {
	if t.synqly.management == nil {
		return nil, "", errors.New("must set synqly.management")
	}

	credential, err := t.synqly.management.Credentials.CreateCredential(ctx, t.synqly.accountId, &mgmt.CreateCredentialRequest{
		Name: "tenable",
		Config: mgmt.NewCredentialConfigFromToken(&mgmt.TokenCredential{
			Secret: "accessKey=ba1a8390271e8f0dabd7996c3ee4c1a743f9eaad0d659a38dc38f04a39ab51f9;secretKey=921073a7d56bde4b73fb32cb89e092fe05a482364355772ff0150ec624b22f7d",
		}),
	})
	if err != nil {
		return nil, "", fmt.Errorf("unable to create credential: %w", err)
	}

	integration, err := t.createVulnerabilityConnector(ctx, "tenable", &mgmt.VulnerabilityConfig{
		CredentialId: credential.Result.Id,
	})
	if err != nil {
		return nil, "", err
	}

	return engineClient.NewClient(
		engineClient.WithAuthToken(integration.Token.Access.Secret),
	), integration.Token.Access.Secret, nil
}

func (t *Tenant) createTicketingConnector(ctx context.Context, providerType string, providerConfig *mgmt.TicketingConfig) (*mgmt.CreateIntegrationResponseResult, error) {
	response, err := t.synqly.management.Integrations.CreateIntegration(ctx, t.synqly.accountId, &mgmt.CreateIntegrationRequest{
		Name:           "Ticketing Connector",
		Category:       "ticketing",
		ProviderType:   providerType,
		ProviderConfig: mgmt.NewProviderConfigFromTicketing(providerConfig),
	})
	if err != nil {
		return nil, err
	}

	return response.Result, nil
}

func (t *Tenant) configureTicketingConnector(ctx context.Context) (*engineClient.Client, string, error) {
	if t.synqly.management == nil {
		return nil, "", errors.New("must set synqly.management")
	}

	credential, err := t.synqly.management.Credentials.CreateCredential(ctx, t.synqly.accountId, &mgmt.CreateCredentialRequest{
		Name: "jira",
		Config: mgmt.NewCredentialConfigFromBasic(&mgmt.BasicCredential{
			Username: "andy@synqly.com",
			Secret:   jiraToken,
		}),
	})
	if err != nil {
		return nil, "", err
	}

	integration, err := t.createTicketingConnector(ctx, "jira", &mgmt.TicketingConfig{
		Endpoint:     mgmt.String("https://synqly.atlassian.net"),
		CredentialId: credential.Result.Id,
	})
	if err != nil {
		return nil, "", err
	}

	return engineClient.NewClient(
		engineClient.WithAuthToken(integration.Token.Access.Secret),
	), integration.Token.Access.Secret, nil
}

// notification creates a ticket in the tenant's ticketing connector for the given vulnerability finding
func (t *Tenant) notification(ctx context.Context, finding *securityfinding.SecurityFinding) error {
	tags := []string{fmt.Sprintf("asset:%s", *finding.Resources[0].Uid), fmt.Sprintf("scan_uid:%s", *finding.Metadata.Uid)}

	// search for existing ticket
	tickets, err := t.synqly.ticketing.Ticketing.QueryTickets(ctx, &engine.QueryTicketsRequest{
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

	tick, err := t.synqly.ticketing.Ticketing.CreateTicket(ctx, &engine.CreateTicketRequest{
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

	if synqlyOrgToken == "" {
		log.Fatal("Must set following environment variables: SYNQLY_ORG_TOKEN")
	}

	ctx := context.Background()

	t, err := NewTenant(ctx, "synqly-demo", "config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	findings, err := t.synqly.vulnerabilities.Vulnerabilities.QueryVulnerabilityFindings(ctx, &engine.QueryFindingsRequest{
		Filter: []*string{
			engine.String("severity[in]Critical,High,Medium"),
			engine.String("finding.last_seen_time[gte]2023-12-10T00:00:00Z"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	consoleLogger.Printf("Found %d security findings from Tenable\n", len(findings.Result))

	for _, finding := range findings.Result {
		if err := t.notification(ctx, finding.SecurityFinding); err != nil {
			log.Fatal(err)
		}
	}
}
