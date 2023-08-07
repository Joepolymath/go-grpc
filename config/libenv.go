package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var envPrefix string

func LoadEnvVariables(prefix string) {
	envPrefix = prefix
	viper.SetEnvPrefix(envPrefix)

	if os.Getenv("GO_ENV") == "local" {
		viper.SetConfigFile(".env.local.yaml")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("Error reading config file: %v", err)
		}
	}
	viper.AutomaticEnv()
}