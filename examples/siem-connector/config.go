package main

import "github.com/spf13/viper"

type Config struct {
	SynqlyOrgToken    string `mapstructure:"SYNQLY_ORG_TOKEN"`
	SplunkHecUrl      string `mapstructure:"SPLUNK_HEC_URL"`
	SplunkHecToken    string `mapstructure:"SPLUNK_HEC_TOKEN"`
	SplunkSearchUrl   string `mapstructure:"SPLUNK_SEARCH_URL"`
	SplunkSearchToken string `mapstructure:"SPLUNK_SEARCH_TOKEN"`
	DurationSeconds   string `mapstructure:"DURATION_SECONDS"`
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
