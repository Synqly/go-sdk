package main

import "github.com/spf13/viper"

type Config struct {
	synqlyOrgToken string `mapstructure:"SYNQLY_ORG_TOKEN"`
	userEmail      string `mapstructure:"USER_EMAIL"`
	url            string `mapstructure:"URL"`
	token          string `mapstructure:"TOKEN"`
	clientId       string `mapstructure:"CLIENt_ID"`
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
