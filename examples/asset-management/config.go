package main

import "github.com/spf13/viper"

type Config struct {
	SynqlyOrgToken    string `mapstructure:"SYNQLY_ORG_TOKEN"`
	ServiceNowURL     string `mapstructure:"SYNQLY_SERVICE_NOW_URL"`
	ServiceNowTokenID string `mapstructure:"SYNQLY_SERVICE_NOW_TOKEN"`
	ArmisURL          string `mapstructure:"SYNQLY_ARMIS_URL"`
	ArmisTokenID      string `mapstructure:"SYNQLY_ARMIS_TOKEN"`
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
