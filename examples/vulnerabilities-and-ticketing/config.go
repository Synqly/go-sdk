package main

import "github.com/spf13/viper"

type Config struct {
	synqlyOrgToken string `mapstructure:"SYNQLY_ORG_TOKEN"`
	jiraUrl        string `mapstructure:"JIRA_URL"`
	jiraUser       string `mapstructure:"JIRA_USER"`
	jiraToken      string `mapstructure:"JIRA_TOKEN"`
	tenableToken   string `mapstructure:"TENABLE_TOKEN"`
	qualysEndpoint string `mapstructure:"QUALYS_ENDPOINT"`
	qualysUser     string `mapstructure:"QUALYS_USER"`
	qualysSecret   string `mapstructure:"QUALYS_SECRET"`
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
