package main

import (
	"log"

	"verifone-soap/pkg/verifone"

	"github.com/spf13/viper"
)

type VerifoneResponse struct {
	Fault FaultResponse `xml:"Fault"`
}

type FaultResponse struct {
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
	Detail      string `xml:"detail"`
}

type Message struct {
	//XMLName struct{} `xml:"Message"`
	//ClientHeader ClientHeader
	MsgType string `xml:"MsgType"`
	MsgData string `xml:"MsgData"`
}

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
	/*
		resp := VerifoneResponse{}
		err = soap.Unmarshal(&resp)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(resp)*/
	//log.Println(string(soap.Body))

}
