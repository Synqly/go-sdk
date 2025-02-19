package main

import (
	"log"

	"github.com/spf13/viper"
	mgmt "github.com/synqly/go-sdk/client/management"
)

type Config struct {
	SynqlyOrgToken  string
	ProviderConfigs []Provider
}

// temp struct to store the env variables
type vars struct {
	SynqlyOrgToken  string `mapstructure:"SYNQLY_ORG_TOKEN"`
	ServicenowURL   string `mapstructure:"SYNQLY_SERVICENOW_URL"`
	ServicenowToken string `mapstructure:"SYNQLY_SERVICENOW_TOKEN"`
	ArmisURL        string `mapstructure:"SYNQLY_ARMIS_URL"`
	ArmisToken      string `mapstructure:"SYNQLY_ARMIS_TOKEN"`
	NozomiURL       string `mapstructure:"SYNQLY_NOZOMI_URL"`
	NozomiUser      string `mapstructure:"SYNQLY_NOZOMI_USER"`
	NozomiToken     string `mapstructure:"SYNQLY_NOZOMI_TOKEN"`
}

type Provider interface {
	Name() string
	ProviderConfig() *mgmt.ProviderConfig
	IsConfigured() bool
}

// ServiceNnow
type servicenowConfig struct {
	URL   string `mapstructure:"SYNQLY_SERVICENOW_URL"`
	Token string `mapstructure:"SYNQLY_SERVICENOW_TOKEN"`
}

func (c *servicenowConfig) ProviderConfig() *mgmt.ProviderConfig {
	return &mgmt.ProviderConfig{
		AssetsServicenow: &mgmt.AssetsServiceNow{
			Url: c.URL,
			Credential: &mgmt.ServiceNowCredential{
				Token: &mgmt.TokenCredential{
					Secret: c.Token,
				},
			},
		},
	}
}

func (c *servicenowConfig) IsConfigured() bool {
	return c.URL != "" && c.Token != ""
}

func (c *servicenowConfig) Name() string {
	return "service-now"
}

// Armis Centrix
type armisConfig struct {
	URL   string `mapstructure:"SYNQLY_ARMIS_URL"`
	Token string `mapstructure:"SYNQLY_ARMIS_TOKEN"`
}

func (c *armisConfig) ProviderConfig() *mgmt.ProviderConfig {
	return &mgmt.ProviderConfig{
		AssetsArmisCentrix: &mgmt.AssetsArmisCentrix{
			Url: c.URL,
			Credential: &mgmt.ArmisCredential{
				Token: &mgmt.TokenCredential{
					Secret: c.Token,
				},
			},
		},
	}
}

func (c *armisConfig) IsConfigured() bool {
	return c.URL != "" && c.Token != ""
}

func (c *armisConfig) Name() string {
	return "armis"
}

// Nozomi Vantage
type nozomiConfig struct {
	URL   string `mapstructure:"SYNQLY_NOZOMI_URL"`
	User  string `mapstructure:"SYNQLY_NOZOMI_USER"`
	Token string `mapstructure:"SYNQLY_NOZOMI_TOKEN"`
}

func (c *nozomiConfig) ProviderConfig() *mgmt.ProviderConfig {
	return &mgmt.ProviderConfig{
		AssetsNozomiVantage: &mgmt.AssetsNozomiVantage{
			Url: c.URL,
			Credential: &mgmt.NozomiVantageCredential{
				Basic: &mgmt.BasicCredential{
					Username: c.User,
					Secret:   c.Token,
				},
			},
		},
	}
}

func (c *nozomiConfig) IsConfigured() bool {
	return c.URL != "" && c.User != "" && c.Token != ""
}

func (c *nozomiConfig) Name() string {
	return "nozomi"
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	var v vars
	err = viper.Unmarshal(&v)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	config.SynqlyOrgToken = v.SynqlyOrgToken
	servicenowConfig := &servicenowConfig{
		URL:   v.ServicenowURL,
		Token: v.ServicenowToken,
	}

	config.ProviderConfigs = append(config.ProviderConfigs, servicenowConfig)
	armisConfig := &armisConfig{
		URL:   v.ArmisURL,
		Token: v.ArmisToken,
	}

	config.ProviderConfigs = append(config.ProviderConfigs, armisConfig)

	nozomiConfig := &nozomiConfig{
		URL:   v.NozomiURL,
		User:  v.NozomiUser,
		Token: v.NozomiToken,
	}

	config.ProviderConfigs = append(config.ProviderConfigs, nozomiConfig)
	return config, nil
}
