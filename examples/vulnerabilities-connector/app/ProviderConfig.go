package app

import (
	"context"
	"fmt"
	"log"

	"examples/common"
	"github.com/synqly/go-sdk/client/engine"
	mgmt "github.com/synqly/go-sdk/client/management"
)

const (
	TENANT_CONFIG_NAME_PREFIX string = "Zenith Systems - "
)

type ProviderConfiguration struct {
	AccountName        string
	IntegrationRequest *mgmt.CreateIntegrationRequest
	Name               string
	Tenant             *common.Tenant
}

func (p *ProviderConfiguration) New(config Configuration, providerConfig mgmt.ProviderConfig, providerName string) error {
	p.AccountName = fmt.Sprintf("%s%s", TENANT_CONFIG_NAME_PREFIX, providerName)
	p.IntegrationRequest = &mgmt.CreateIntegrationRequest{
		Fullname:       engine.String(fmt.Sprintf("Vulnerability %s Scanner", providerName)),
		ProviderConfig: &providerConfig,
	}

	p.Name = providerName

	// Tenant
	configFilePath := fmt.Sprintf("./tenant_config_files/tenant_store_%s.yaml", p.Name)

	ctx := context.Background()
	t, err := common.NewTenant(ctx, p.AccountName, configFilePath, config.SynqlyOrgToken, map[mgmt.CategoryId]*mgmt.CreateIntegrationRequest{
		mgmt.CategoryIdVulnerabilities: p.IntegrationRequest,
	})
	if err != nil {
		log.Printf("Error creating the Tenant for %s: %s", p.Name, err.Error())
		return err
	}

	p.Tenant = t

	return nil
}
