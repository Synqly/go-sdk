package common

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	engineClient "github.com/synqly/go-sdk/client/engine/client"
	mgmt "github.com/synqly/go-sdk/client/management"
	mgmtClient "github.com/synqly/go-sdk/client/management/client"
)

var consoleLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

// Tenant represents a customer of yours. It has its own account in Synqly, and this struct
// keeps the information you need to interact with Synqly on behalf of the customer.
type Tenant struct {
	ID     string
	Synqly SynqlyRuntime
}

// SynqlyRuntime keeps the clients needed to use the Synqly management, identity and ticketing APIs
type SynqlyRuntime struct {
	accountId     string
	management    *mgmtClient.Client
	EngineClients map[mgmt.CategoryId]*engineClient.Client
}

// TenantConfig is the configuration you need to save for your tenants to re-connect to Synqly
type TenantConfig struct {
	SynqlyConfig SynqlyConfig `yaml:"synqly_config"`
}

type SynqlyConfig struct {
	AccountID string                     `yaml:"account_id"`
	Tokens    map[mgmt.CategoryId]string `yaml:"tokens"`
}

// NewTenant creates a tenant and configures it with Synqly. If this tenant does not
// have a saved config file, then a new Synqly Account will be created with a identity
// connector. If the tenant already has a config file, then the existing
// Synqly Account and connectors will be used.
func NewTenant(ctx context.Context, id, configFilePath, synqlyOrgToken string, integrations map[mgmt.CategoryId]*mgmt.CreateIntegrationRequest) (*Tenant, error) {
	tenant := &Tenant{}
	config := &TenantConfig{}

	configFile, err := os.Open(configFilePath)
	if errors.Is(err, os.ErrNotExist) {
		configFile, err = os.Create(configFilePath)
	}
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	if err := yaml.NewDecoder(configFile).Decode(&config); err != nil {
		config = &TenantConfig{}
	}

	if config.SynqlyConfig.Tokens == nil {
		config.SynqlyConfig.Tokens = make(map[mgmt.CategoryId]string)
	}

	if tenant.Synqly.EngineClients == nil {
		tenant.Synqly.EngineClients = make(map[mgmt.CategoryId]*engineClient.Client)
	}

	tenant.Synqly.management = mgmtClient.NewClient(
		mgmtClient.WithAuthToken(synqlyOrgToken),
	)

	if err := tenant.Init(ctx, id, config, integrations); err != nil {
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
func (t *Tenant) Init(ctx context.Context, tenantID string, config *TenantConfig, integrations map[mgmt.CategoryId]*mgmt.CreateIntegrationRequest) error {
	accountId := config.SynqlyConfig.AccountID
	if accountId == "" {
		account, err := t.Synqly.management.Accounts.Create(ctx, &mgmt.CreateAccountRequest{
			Fullname: mgmt.String(tenantID),
		})
		if err != nil {
			return fmt.Errorf("init failed, unable to create account: %w", err)
		}
		accountId = account.Result.Account.Id
		consoleLogger.Printf("Created new account: %s", accountId)
	}
	config.SynqlyConfig.AccountID = accountId
	t.Synqly.accountId = accountId

	consoleLogger.Printf("Using account with id: %s\n", accountId)

	var err error
	var token string

	for categoryID, integration := range integrations {
		consoleLogger.Printf("Configuring %s connector\n", integration.ProviderConfig.Type)

		if config.SynqlyConfig.Tokens[categoryID] == "" {
			t.Synqly.EngineClients[categoryID], token, err = t.ConfigureConnector(ctx, integration)
			if err != nil {
				return fmt.Errorf("init failed, unable to configure %s connector: %w", categoryID, err)
			}
			config.SynqlyConfig.Tokens[categoryID] = token
		} else {
			t.Synqly.EngineClients[categoryID] = engineClient.NewClient(
				engineClient.WithAuthToken(config.SynqlyConfig.Tokens[categoryID]),
			)
		}
	}

	return nil
}

func (t *Tenant) ConfigureConnector(ctx context.Context, integration *mgmt.CreateIntegrationRequest) (*engineClient.Client, string, error) {
	if t.Synqly.management == nil {
		return nil, "", errors.New("must set synqly.management")
	}

	resp, err := t.Synqly.management.Integrations.Create(ctx, t.Synqly.accountId, integration)
	if err != nil {
		return nil, "", fmt.Errorf("unable to create integration: %w", err)
	}

	return engineClient.NewClient(
		engineClient.WithAuthToken(resp.Result.Token.Access.Secret),
	), resp.Result.Token.Access.Secret, nil
}
