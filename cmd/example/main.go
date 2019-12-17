package main

import (
	"log"

	"github.com/nazwa/verifone-soap/pkg/verifone"
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
	}, "")

	session, err := client.BeginSession(viper.GetString("verifone.Redirect"), true)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("session: %s, %s, %s\n\r", session.SessionGUID, session.SessionPasscode, session.ProcessingDB)
	}

	token, err := client.RegisterToken(session.SessionGUID, "", "", "20022019", true, false, false)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("token:", token.TokenId, token.ErrorCode, token.ErrorDescription)
	}

}
