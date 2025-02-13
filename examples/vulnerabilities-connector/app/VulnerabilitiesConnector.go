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
	Provider                VulnerabilitiesConnectorProvider
	Tenant                  *common.Tenant
	VulnerabilitiesProvider *mgmt.CreateIntegrationRequest
}

func (v *VulnerabilitiesConnector) Run() {
	v.Logger.Printf("Using %s as vulnerability provider\n", v.VulnerabilitiesProvider.ProviderConfig.Type)

	ctx := context.Background()

	client := v.Tenant.Synqly.EngineClients["vulnerabilities"]

	v.QueryAssets(ctx, client)
	v.QueryFindings(ctx, client)
	v.QueryScans(ctx, client)
}

func (v *VulnerabilitiesConnector) Initialize(provider VulnerabilitiesConnectorProvider) {
	// Config
	v.Config = InitConfig()

	// Logger
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	v.Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	// Vulnerabilities Provider
	v.Provider = provider

	vulnProvider, err := v.buildVulnerabilityProviderConfig(provider)
	if err != nil {
		log.Fatal(err)
	}
	v.VulnerabilitiesProvider = vulnProvider

	tenantId := fmt.Sprintf("Zenith Systems - %s", v.Provider)
	configFilePath := fmt.Sprintf("tenant_store_%s.yaml", v.Provider)

	ctx := context.Background()
	t, err := common.NewTenant(ctx, tenantId, configFilePath, v.Config.SynqlyOrgToken, map[mgmt.CategoryId]*mgmt.CreateIntegrationRequest{
		mgmt.CategoryIdVulnerabilities: vulnProvider,
	})
	if err != nil {
		log.Fatal(err)
	}
	v.Tenant = t

	v.Logger.Printf("Example initialized successfully")
}

func (v *VulnerabilitiesConnector) QueryAssets(ctx context.Context, client *engineClient.Client) {
	filter := []*string{}

	if v.Provider != NUCLEUS {
		filter = append(filter, engine.String("device.last_seen_time[gte]2024-09-01T00:00:00Z"))
	}

	assets, err := client.Vulnerabilities.QueryAssets(ctx, &engine.QueryAssetsRequest{
		Filter: filter,
	})
	if err != nil {
		log.Fatal(err)
	}
	v.Logger.Printf("Found %d security assets from %s\n", len(assets.Result), v.VulnerabilitiesProvider.ProviderConfig.Type)
}

func (v *VulnerabilitiesConnector) QueryFindings(ctx context.Context, client *engineClient.Client) {
	filter := []*string{
		engine.String("severity[in]Critical"),
	}

	switch v.Provider {
	case CROUD_STRIKE:
		filter = append(filter, engine.String("finding.last_seen_time[gte]2024-09-01T00:00:00Z"))
	case NUCLEUS:
	case QUALYS:
		filter = append(filter, engine.String("finding.last_seen_time[gte]2024-09-01T00:00:00Z"))
	case RAPID7:
		filter = append(filter, engine.String("finding.last_seen_time[gte]2024-09-01T00:00:00Z"))
	case TANIUM:
		filter = append(filter, engine.String("finding.last_seen_time[gte]2024-09-01T00:00:00Z"))
	case TENABLE:
		filter = append(filter, engine.String("state[in]New"))
		filter = append(filter, engine.String("finding.last_seen_time[gte]2024-09-01T00:00:00Z"))
	default:
		log.Fatalf("Please provide a valid provider: %s, %s, %s, %s, %s or %s", CROUD_STRIKE, NUCLEUS, QUALYS, RAPID7, TANIUM, TENABLE)
	}

	findings, err := client.Vulnerabilities.QueryFindings(ctx, &engine.QueryFindingsRequest{
		Filter: filter,
	})
	if err != nil {
		log.Fatal(err)
	}
	v.Logger.Printf("Found %d security findings from %s\n", len(findings.Result), v.VulnerabilitiesProvider.ProviderConfig.Type)
}

func (v *VulnerabilitiesConnector) QueryScans(ctx context.Context, client *engineClient.Client) {
	var e *engine.NotImplementedError

	assets, err := client.Vulnerabilities.QueryScans(ctx, &engine.QueryScansRequest{})
	if err != nil && !errors.As(err, &e) {
		log.Fatal(err)
	}

	if errors.As(err, &e) {
		v.Logger.Println("Not implemented")
	} else {
		v.Logger.Printf("Found %d scans from %s\n", len(assets.Result), v.VulnerabilitiesProvider.ProviderConfig.Type)
	}
}

func (v *VulnerabilitiesConnector) buildVulnerabilityProviderConfig(provider VulnerabilitiesConnectorProvider) (*mgmt.CreateIntegrationRequest, error) {
	// CroudStrike
	if provider == CROUD_STRIKE {
		if v.Config.CroudStrikeClientID != "" && v.Config.CroudStrikeClientSecret != "" && v.Config.CroudStrikeUrl != "" {
			return &mgmt.CreateIntegrationRequest{
				Fullname: engine.String(fmt.Sprintf("Vulnerability %s Scanner", v.Provider)),
				ProviderConfig: &mgmt.ProviderConfig{
					Type: string(v.Provider),
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
				Fullname: engine.String(fmt.Sprintf("Vulnerability %s Scanner", v.Provider)),
				ProviderConfig: &mgmt.ProviderConfig{
					Type: string(v.Provider),
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
				Fullname: engine.String(fmt.Sprintf("Vulnerability %s Scanner", v.Provider)),
				ProviderConfig: &mgmt.ProviderConfig{
					Type: string(v.Provider),
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
				Fullname: engine.String(fmt.Sprintf("Vulnerability %s Scanner", v.Provider)),
				ProviderConfig: &mgmt.ProviderConfig{
					Type: string(v.Provider),
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
				Fullname: engine.String(fmt.Sprintf("Vulnerability %s Scanner", v.Provider)),
				ProviderConfig: &mgmt.ProviderConfig{
					Type: string(v.Provider),
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
					Type: string(v.Provider),
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
