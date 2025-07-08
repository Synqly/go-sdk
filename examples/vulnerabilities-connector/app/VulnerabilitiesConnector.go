package app

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/synqly/go-sdk/client/engine"
	engineClient "github.com/synqly/go-sdk/client/engine/client"
	mgmt "github.com/synqly/go-sdk/client/management"
	mgmtClient "github.com/synqly/go-sdk/client/management/client"
)

type VulnerabilitiesConnector struct {
	Config    *Configuration
	Logger    *log.Logger
	Providers []ProviderConfiguration
}

func NewVulnerabilitiesConnector() *VulnerabilitiesConnector {
	vulnerabilitiesConnector := VulnerabilitiesConnector{}
	vulnerabilitiesConnector.Initialize()

	return &vulnerabilitiesConnector
}

func (v *VulnerabilitiesConnector) CleanUp(cleanUp bool) {
	if !cleanUp {
		return
	}

	v.Logger.Println("CleanUp started...")

	client := mgmtClient.NewClient(mgmtClient.WithToken(v.Config.SynqlyOrgToken))

	if client == nil {
		return
	}

	ctx := context.Background()

	allAccounts, err := client.Accounts.List(ctx, &mgmt.ListAccountsRequest{
		Filter: []*string{engine.String("name[like]zenith-systems - *")},
	})
	if err != nil {
		v.Logger.Printf("Error checking account IDs: %s\n", err.Error())
	}

	getAccountId := func(fullName string) (string, error) {
		for _, account := range allAccounts.Result {
			if account.Fullname == fullName {
				return account.Id, nil
			}
		}

		return "", fmt.Errorf("account ID for '%s' not found", fullName)
	}

	for _, p := range v.Providers {
		accountID, err := getAccountId(p.AccountName)
		if err != nil {
			v.Logger.Printf("Error getting account '%s' ID, please delete account manually: %s\n", p.AccountName, err.Error())
			continue
		}

		err = client.Accounts.Delete(ctx, accountID)
		if err != nil {
			v.Logger.Printf("Error deleting '%s' account, please do it manually: %s\n", p.AccountName, err.Error())
			continue
		}

		v.Logger.Printf("Accound '%s' deleted...\n", p.AccountName)
	}

	err = os.RemoveAll("./tenant_config_files")
	if err != nil {
		v.Logger.Printf("Error deleting 'tenant_config_files' folder, please do it manually: %s\n", err.Error())
	}

	v.Logger.Println("CleanUp finished successfully")
}

func (v *VulnerabilitiesConnector) Initialize() {
	// Tenants Config Folder
	_, err := os.Stat("./tenant_config_files")
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		os.Mkdir("./tenant_config_files", 0700)
	}

	// Init
	v.Config = InitConfig()
	v.Providers = []ProviderConfiguration{}

	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	v.Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	// Providers
	v.Providers = v.buildVulnerabilityProvidersConfig()

	v.Logger.Println("Example initialized...")
}

func (v *VulnerabilitiesConnector) QueryAssets(ctx context.Context, client *engineClient.Client, provider *mgmt.CreateIntegrationRequest) {
	filter := []*string{}

	assets, err := client.Vulnerabilities.QueryAssets(ctx, &engine.QueryAssetsRequest{
		Filter: filter,
	})
	if err != nil {
		v.Logger.Printf("Error executing QueryAssets for %s provider: %s\n", provider, err.Error())
		return
	}
	v.Logger.Printf("Found %d security assets from %s\n", len(assets.Result), provider.ProviderConfig.Type)
}

func (v *VulnerabilitiesConnector) QueryFindings(ctx context.Context, client *engineClient.Client, provider *mgmt.CreateIntegrationRequest) {
	filter := []*string{
		engine.String("severity[in]Critical"),
	}

	findings, err := client.Vulnerabilities.QueryFindings(ctx, &engine.QueryFindingsRequest{
		Filter: filter,
	})
	if err != nil {
		v.Logger.Printf("Error executing QueryFindings for %s provider: %s\n", provider, err.Error())
		return
	}
	v.Logger.Printf("Found %d security findings from %s\n", len(findings.Result), provider.ProviderConfig.Type)
}

func (v *VulnerabilitiesConnector) QueryScans(ctx context.Context, client *engineClient.Client, provider *mgmt.CreateIntegrationRequest) {
	var e *engine.NotImplementedError

	assets, err := client.Vulnerabilities.QueryScans(ctx, &engine.QueryScansRequest{})
	if err != nil && !errors.As(err, &e) {
		v.Logger.Printf("Error executing QueryScans for %s provider: %s\n", provider, err.Error())
		return
	} else if errors.As(err, &e) {
		// not all providers support the QueryScans operation. If an operation is not supported, the API will return
		// a NotImplementedError so we can check for that error and handle it specifically

		v.Logger.Printf("Query scans not implemented for %s\n", provider.ProviderConfig.Type)
		return
	}

	v.Logger.Printf("Found %d scans from %s\n", len(assets.Result), provider.ProviderConfig.Type)
}

func (v *VulnerabilitiesConnector) Start(cleanUp bool) {
	defer v.CleanUp(cleanUp)

	v.Logger.Printf("Using %d vulnerability providers\n", len(v.Providers))

	ctx := context.Background()

	for _, provider := range v.Providers {
		client := provider.Tenant.Synqly.EngineClients["vulnerabilities"]

		v.QueryAssets(ctx, client, provider.IntegrationRequest)
		v.QueryFindings(ctx, client, provider.IntegrationRequest)
		v.QueryScans(ctx, client, provider.IntegrationRequest)
	}
}

func (v *VulnerabilitiesConnector) buildVulnerabilityProvidersConfig() []ProviderConfiguration {
	providers := []ProviderConfiguration{}

	// CroudStrike
	if v.Config.CroudStrikeClientID != "" && v.Config.CroudStrikeClientSecret != "" && v.Config.CroudStrikeUrl != "" {
		croudStrikeConfig := ProviderConfiguration{}
		if err := croudStrikeConfig.New(*v.Config, mgmt.ProviderConfig{
			Type: "CroudStrike",
			VulnerabilitiesCrowdstrike: &mgmt.VulnerabilitiesCrowdStrike{
				Credential: &mgmt.CrowdStrikeCredential{
					OAuthClient: &mgmt.OAuthClientCredential{
						ClientId:     v.Config.CroudStrikeClientID,
						ClientSecret: v.Config.CroudStrikeClientSecret,
					},
				},
				Url: &v.Config.CroudStrikeUrl,
			},
		}, "CroudStrike"); err != nil {
			v.Logger.Printf("There was an error initializing CroudStrike configuration: %s\n", err.Error())
		} else {
			providers = append(providers, croudStrikeConfig)
		}
	} else {
		v.Logger.Println("Ignoring CroudStrike provider, missing configuration variables...")
	}

	// Nucleus
	if v.Config.NucleusProjectID != "" && v.Config.NucleusToken != "" && v.Config.NucleusUrl != "" {
		nucleusConfig := ProviderConfiguration{}
		if err := nucleusConfig.New(*v.Config, mgmt.ProviderConfig{
			Type: "Nucleus",
			VulnerabilitiesNucleus: &mgmt.VulnerabilitiesNucleus{
				Credential: &mgmt.NucleusCredential{
					Token: &mgmt.TokenCredential{
						Secret: v.Config.NucleusToken,
					},
				},
				ProjectId: v.Config.NucleusProjectID,
				Url:       v.Config.NucleusUrl,
			},
		}, "Nucleus"); err != nil {
			v.Logger.Printf("There was an error initializing Nucleus configuration: %s\n", err.Error())
		} else {
			providers = append(providers, nucleusConfig)
		}
	} else {
		v.Logger.Println("Ignoring Nucleus provider, missing configuration variables...")
	}

	// Qualys
	if v.Config.QualysSecret != "" && v.Config.QualysUrl != "" && v.Config.QualysUser != "" {
		qualysConfig := ProviderConfiguration{}
		if err := qualysConfig.New(*v.Config, mgmt.ProviderConfig{
			Type: "Qualys",
			VulnerabilitiesQualysCloud: &mgmt.VulnerabilitiesQualysCloud{
				Credential: &mgmt.QualysCloudCredential{
					Basic: &mgmt.BasicCredential{
						Username: v.Config.QualysUser,
						Secret:   v.Config.QualysSecret,
					},
				},
				Url: v.Config.QualysUrl,
			},
		}, "Qualys"); err != nil {
			v.Logger.Printf("There was an error initializing Qualys configuration: %s\n", err.Error())
		} else {
			providers = append(providers, qualysConfig)
		}
	} else {
		v.Logger.Println("Ignoring Qualys provider, missing configuration variables...")
	}

	// Rapid7
	if v.Config.Rapid7Token != "" && v.Config.Rapid7Url != "" {
		rapid7Config := ProviderConfiguration{}
		if err := rapid7Config.New(*v.Config, mgmt.ProviderConfig{
			Type: "Rapid7",
			VulnerabilitiesRapid7InsightCloud: &mgmt.VulnerabilitiesRapid7InsightCloud{
				Url: v.Config.Rapid7Url,
				Credential: &mgmt.Rapid7InsightCloudCredential{
					Token: &mgmt.TokenCredential{
						Secret: v.Config.Rapid7Token,
					},
				},
			},
		}, "Rapid7"); err != nil {
			v.Logger.Printf("There was an error initializing Rapid7 configuration: %s\n", err.Error())
		} else {
			providers = append(providers, rapid7Config)
		}
	} else {
		v.Logger.Println("Ignoring Rapid7 provider, missing configuration variables...")
	}

	// Tanium
	if v.Config.TaniumSecret != "" && v.Config.TaniumUrl != "" {
		taniumConfig := ProviderConfiguration{}
		if err := taniumConfig.New(*v.Config, mgmt.ProviderConfig{
			Type: "Tanium",
			VulnerabilitiesTaniumCloud: &mgmt.VulnerabilitiesTaniumCloud{
				Credential: &mgmt.TaniumCloudCredential{
					Token: &mgmt.TokenCredential{
						Secret: v.Config.TaniumSecret,
					},
				},
				Url: v.Config.TaniumUrl,
			},
		}, "Tanium"); err != nil {
			v.Logger.Printf("There was an error initializing Tanium configuration: %s\n", err.Error())
		} else {
			providers = append(providers, taniumConfig)
		}
	} else {
		v.Logger.Println("Ignoring Tanium provider, missing configuration variables...")
	}

	// Tenable
	if v.Config.TenableToken != "" && v.Config.TenableUrl != "" {
		tenableConfig := ProviderConfiguration{}
		if err := tenableConfig.New(*v.Config, mgmt.ProviderConfig{
			Type: "Tenable",
			VulnerabilitiesTenableCloud: &mgmt.VulnerabilitiesTenableCloud{
				Credential: &mgmt.TenableCloudCredential{
					Token: &mgmt.TokenCredential{
						Secret: v.Config.TenableToken,
					},
				},
				Url: &v.Config.TenableUrl,
			},
		}, "Tenable"); err != nil {
			v.Logger.Printf("There was an error initializing Tenable configuration: %s\n", err.Error())
		} else {
			providers = append(providers, tenableConfig)
		}
	} else {
		v.Logger.Println("Ignoring Tenable provider, missing configuration variables...")
	}

	return providers
}
