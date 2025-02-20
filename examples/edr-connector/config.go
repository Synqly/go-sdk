package main

import "github.com/spf13/viper"

type Config struct {
	synqlyOrgToken string `mapstructure:"SYNQLY_ORG_TOKEN"`
	edrURL         string `mapstructure:"EDR_URL"`
	edrToken       string `mapstructure:"EDR_TOKEN"`
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
