package app

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	CroudStrikeClientID     string `mapstructure:"SYNQLY_CROUDSTRIKE_CLIENT_ID"`
	CroudStrikeClientSecret string `mapstructure:"SYNQLY_CROUDSTRIKE_CLIENT_SECRET"`
	CroudStrikeUrl          string `mapstructure:"SYNQLY_CROUDSTRIKE_URL"`

	NucleusProjectID string `mapstructure:"SYNQLY_NUCLEUS_PROJECT_ID"`
	NucleusToken     string `mapstructure:"SYNQLY_NUCLEUS_TOKEN"`
	NucleusUrl       string `mapstructure:"SYNQLY_NUCLEUS_URL"`

	QualysSecret string `mapstructure:"SYNQLY_QUALYS_SECRET"`
	QualysUrl    string `mapstructure:"SYNQLY_QUALYS_URL"`
	QualysUser   string `mapstructure:"SYNQLY_QUALYS_USER"`

	Rapid7Token string `mapstructure:"SYNQLY_RAPID7_TOKEN"`
	Rapid7Url   string `mapstructure:"SYNQLY_RAPID7_URL"`

	SynqlyOrgToken string `mapstructure:"SYNQLY_ORG_TOKEN"`

	TaniumSecret string `mapstructure:"SYNQLY_TANIUM_SECRET"`
	TaniumUrl    string `mapstructure:"SYNQLY_TANIUM_URL"`

	TenableToken string `mapstructure:"SYNQLY_TENABLE_TOKEN"`
	TenableUrl   string `mapstructure:"SYNQLY_TENABLE_URL"`
}

func InitConfig() *Configuration {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("error reading config")
	}

	configuration := Configuration{}

	viper.Unmarshal(&configuration)

	return &configuration
}
