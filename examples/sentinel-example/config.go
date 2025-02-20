package main

import "github.com/spf13/viper"

type Config struct {
	synqlyOrgToken    string `mapstructure:"SYNQLY_ORG_TOKEN"`
	azureTenantId     string `mapstructure:"AZURE_TENANT_ID"`
	azureClientId     string `mapstructure:"AZURE_CLIENT_ID"`
	azureClientSecret string `mapstructure:"AZURE_CLIENT_SECRET"`
	dceURL            string `mapstructure:"AZURE_DCE_URL"`
	dcrId             string `mapstructure:"AZURE_DCR_ID"`
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

	err = viper.Unmarshal(&config)
	return
}
