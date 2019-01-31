package main

import (
	"log"
	"verifone-soap/pkg/verifone"

	"github.com/spf13/viper"
)

func main() {
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Config file not loaded: %s", err)
	}

	client := verifone.NewClient(verifone.Config{
		SystemID:   viper.GetString("verifone.SystemID"),
		SystemGUID: viper.GetString("verifone.SystemGUID"),
		Passcode:   viper.GetString("verifone.Passcode"),
		Url:        viper.GetString("verifone.Url"),
	})

	token, err := client.BeginSession()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("token:", token)
	}
}
