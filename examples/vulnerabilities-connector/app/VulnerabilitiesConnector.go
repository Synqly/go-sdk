package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/synqly/go-sdk/client/engine"
	"github.com/synqly/go-sdk/examples/common"

	engineClient "github.com/synqly/go-sdk/client/engine/client"
	mgmt "github.com/synqly/go-sdk/client/management"
)

type VulnerabilitiesConnector struct {
	Config                  *Configuration
	Logger                  *log.Logger
	Providers               []VulnerabilitiesConnectorProvider
	Tenants                 map[VulnerabilitiesConnectorProvider]*common.Tenant
	VulnerabilitiesProvider map[VulnerabilitiesConnectorProvider]*mgmt.CreateIntegrationRequest
}

func (v *VulnerabilitiesConnector) Run() {
	v.Logger.Printf("Using %+v as vulnerability provider\n", v.Providers)

	ctx := context.Background()

	for _, provider := range v.Providers {
		if tenant, ok := v.Tenants[provider]; ok {
			client := tenant.Synqly.EngineClients["vulnerabilities"]

			if vulnerabilityProvider, ok := v.VulnerabilitiesProvider[provider]; ok {
				v.QueryAssets(ctx, client, vulnerabilityProvider)
				v.QueryFindings(ctx, client, vulnerabilityProvider)
				v.QueryScans(ctx, client, vulnerabilityProvider)
			}
		}
	}
}

func (v *VulnerabilitiesConnector) Initialize(providers []VulnerabilitiesConnectorProvider) {
	// Config
	v.Config = InitConfig()
	v.Tenants = map[VulnerabilitiesConnectorProvider]*common.Tenant{}
	v.VulnerabilitiesProvider = map[VulnerabilitiesConnectorProvider]*mgmt.CreateIntegrationRequest{}

	// Logger
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	v.Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	// Vulnerabilities Provider
	v.Providers = providers

	for _, provider := range v.Providers {
		vulnProvider, err := v.buildVulnerabilityProviderConfig(provider)
		if err != nil {
			log.Printf("Error building %s config: %s", provider, err.Error())
			continue
		}

		tenantId := fmt.Sprintf("Zenith Systems - %s", provider)
		configFilePath := fmt.Sprintf("tenant_store_%s.yaml", provider)

		ctx := context.Background()
		t, err := common.NewTenant(ctx, tenantId, configFilePath, v.Config.SynqlyOrgToken, map[mgmt.CategoryId]*mgmt.CreateIntegrationRequest{
			mgmt.CategoryIdVulnerabilities: vulnProvider,
		})
		if err != nil {
			log.Printf("Error creating the Tenant for %s: %s", provider, err.Error())
			continue
		}

		v.VulnerabilitiesProvider[provider] = vulnProvider
		v.Tenants[provider] = t
	}

	v.Logger.Printf("Example initialized successfully")
}

func (v *VulnerabilitiesConnector) QueryAssets(ctx context.Context, client *engineClient.Client, provider *mgmt.CreateIntegrationRequest) {
	filter := []*string{}

	if provider.ProviderConfig.Type != string(NUCLEUS) {
		filter = append(filter, engine.String("device.last_seen_time[gte]2024-09-01T00:00:00Z"))
	}

	assets, err := client.Vulnerabilities.QueryAssets(ctx, &engine.QueryAssetsRequest{
		Filter: filter,
	})
	if err != nil {
		log.Printf("Error executing QueryAssets for %s provider: %s", provider, err.Error())
		return
	}
	v.Logger.Printf("Found %d security assets from %s\n", len(assets.Result), provider.ProviderConfig.Type)
}

func (v *VulnerabilitiesConnector) QueryFindings(ctx context.Context, client *engineClient.Client, provider *mgmt.CreateIntegrationRequest) {
	filter := []*string{
		engine.String("severity[in]Critical"),
	}

	if provider.ProviderConfig.Type != string(NUCLEUS) {
		filter = append(filter, engine.String("finding.last_seen_time[gte]2024-09-01T00:00:00Z"))
	}

	if provider.ProviderConfig.Type == string(TENABLE) {
		filter = append(filter, engine.String("state[in]New"))
	}

	findings, err := client.Vulnerabilities.QueryFindings(ctx, &engine.QueryFindingsRequest{
		Filter: filter,
	})
	if err != nil {
		log.Printf("Error executing QueryFindings for %s provider: %s", provider, err.Error())
		return
	}
	v.Logger.Printf("Found %d security findings from %s\n", len(findings.Result), provider.ProviderConfig.Type)
}

func (v *VulnerabilitiesConnector) QueryScans(ctx context.Context, client *engineClient.Client, provider *mgmt.CreateIntegrationRequest) {
	var e *engine.NotImplementedError

	assets, err := client.Vulnerabilities.QueryScans(ctx, &engine.QueryScansRequest{})
	if err != nil && !errors.As(err, &e) {
		log.Printf("Error executing QueryScans for %s provider: %s", provider, err.Error())
		return
	}

	if errors.As(err, &e) {
		v.Logger.Printf("Query scans not implemented for %s\n", provider.ProviderConfig.Type)
	} else {
		v.Logger.Printf("Found %d scans from %s\n", len(assets.Result), provider.ProviderConfig.Type)
	}
}

func (v *VulnerabilitiesConnector) buildVulnerabilityProviderConfig(provider VulnerabilitiesConnectorProvider) (*mgmt.CreateIntegrationRequest, error) {
	// CroudStrike
	if provider == CROUD_STRIKE {
		if v.Config.CroudStrikeClientID != "" && v.Config.CroudStrikeClientSecret != "" && v.Config.CroudStrikeUrl != "" {
			return &mgmt.CreateIntegrationRequest{
				Fullname: engine.String(fmt.Sprintf("Vulnerability %s Scanner", provider)),
				ProviderConfig: &mgmt.ProviderConfig{
					Type: string(provider),
					VulnerabilitiesCrowdstrike: &mgmt.VulnerabilitiesCrowdStrike{
						Credential: &mgmt.CrowdStrikeCredential{
							OAuthClient: &mgmt.OAuthClientCredential{
								ClientId:     v.Config.CroudStrikeClientID,
								ClientSecret: v.Config.CroudStrikeClientSecret,
							},
						},
						Url: &v.Config.CroudStrikeUrl,
					},
				},
			}, nil
		} else {
			return nil, fmt.Errorf("missing %s provider config", provider)
		}
	}

	// Nucleus
	if provider == NUCLEUS {
		if v.Config.NucleusProjectID != "" && v.Config.NucleusToken != "" && v.Config.NucleusUrl != "" {
			return &mgmt.CreateIntegrationRequest{
				Fullname: engine.String(fmt.Sprintf("Vulnerability %s Scanner", provider)),
				ProviderConfig: &mgmt.ProviderConfig{
					Type: string(provider),
					VulnerabilitiesNucleus: &mgmt.VulnerabilitiesNucleus{
						Credential: &mgmt.NucleusCredential{
							Token: &mgmt.TokenCredential{
								Secret: v.Config.NucleusToken,
							},
						},
						ProjectId: v.Config.NucleusProjectID,
						Url:       v.Config.NucleusUrl,
					},
				},
			}, nil
		} else {
			return nil, fmt.Errorf("missing %s provider config", provider)
		}
	}

	// Qualys
	if provider == QUALYS {
		if v.Config.QualysSecret != "" && v.Config.QualysUrl != "" && v.Config.QualysUser != "" {
			return &mgmt.CreateIntegrationRequest{
				Fullname: engine.String(fmt.Sprintf("Vulnerability %s Scanner", provider)),
				ProviderConfig: &mgmt.ProviderConfig{
					Type: string(provider),
					VulnerabilitiesQualysCloud: &mgmt.VulnerabilitiesQualysCloud{
						Credential: &mgmt.QualysCloudCredential{
							Basic: &mgmt.BasicCredential{
								Username: v.Config.QualysUser,
								Secret:   v.Config.QualysSecret,
							},
						},
						Url: v.Config.QualysUrl,
					},
				},
			}, nil
		} else {
			return nil, fmt.Errorf("missing %s provider config", provider)
		}
	}

	// Rapid7
	if provider == RAPID7 {
		if v.Config.Rapid7Token != "" && v.Config.Rapid7Url != "" {
			return &mgmt.CreateIntegrationRequest{
				Fullname: engine.String(fmt.Sprintf("Vulnerability %s Scanner", provider)),
				ProviderConfig: &mgmt.ProviderConfig{
					Type: string(provider),
					VulnerabilitiesRapid7InsightCloud: &mgmt.VulnerabilitiesRapid7InsightCloud{
						Url: v.Config.Rapid7Url,
						Credential: &mgmt.Rapid7InsightCloudCredential{
							Token: &mgmt.TokenCredential{
								Secret: v.Config.Rapid7Token,
							},
						},
					},
				},
			}, nil
		} else {
			return nil, fmt.Errorf("missing %s provider config", provider)
		}

	}

	// Tanium
	if provider == TANIUM {
		if v.Config.TaniumSecret != "" && v.Config.TaniumUrl != "" {
			return &mgmt.CreateIntegrationRequest{
				Fullname: engine.String(fmt.Sprintf("Vulnerability %s Scanner", provider)),
				ProviderConfig: &mgmt.ProviderConfig{
					Type: string(provider),
					VulnerabilitiesTaniumCloud: &mgmt.VulnerabilitiesTaniumCloud{
						Credential: &mgmt.TaniumCloudCredential{
							Token: &mgmt.TokenCredential{
								Secret: v.Config.TaniumSecret,
							},
						},
						Url: v.Config.TaniumUrl,
					},
				},
			}, nil
		} else {
			return nil, fmt.Errorf("missing %s provider config", provider)
		}
	}

	// Tenable
	if provider == TENABLE {
		if v.Config.TenableToken != "" && v.Config.TenableUrl != "" {
			return &mgmt.CreateIntegrationRequest{
				Fullname: engine.String("Vulnerability Tenable Scanner"),
				ProviderConfig: &mgmt.ProviderConfig{
					Type: string(provider),
					VulnerabilitiesTenableCloud: &mgmt.VulnerabilitiesTenableCloud{
						Credential: &mgmt.TenableCloudCredential{
							Token: &mgmt.TokenCredential{
								Secret: v.Config.TenableToken,
							},
						},
						Url: &v.Config.TenableUrl,
					},
				},
			}, nil
		} else {
			return nil, fmt.Errorf("missing %s provider config", provider)
		}
	}

	return nil, fmt.Errorf("configuration for %s not supported", provider)
}
