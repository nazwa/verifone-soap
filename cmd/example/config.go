package main

import (
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("toml")
	viper.AddConfigPath(".") // optionally look for config in the working directory

	// Convert . to _ in ENV names
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
}
