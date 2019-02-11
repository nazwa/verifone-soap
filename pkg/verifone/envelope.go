package verifone

import (
	"encoding/xml"
)

type MsgData struct {
	XMLName xml.Name `xml:"MsgData"`
	Content []byte   `xml:",cdata"`
}

type Message struct {
	XMLName      xml.Name    `xml:"Message"`
	ClientHeader interface{} `xml:"ClientHeader"`
	MsgType      string      `xml:"MsgType"`
	MsgData      MsgData     `xml:"MsgData"`
}

type ProcessMsg struct {
	XMLName xml.Name `xml:"ProcessMsg"`
	Ns      string   `xml:"xmlns,attr"`
	Message interface{}
}

type ClientHeaderResponse struct {
	XMLName      xml.Name `xml:"ClientHeader"`
	ProcessingDB string   `xml:"ProcessingDB"`
	SendAttempt  int64    `xml:"SendAttempt"`
}

type ProcessMsgResult struct {
	ClientHeader *ClientHeaderResponse
	MsgType      string
	MsgData      []byte
}

type ProcessMsgResponse struct {
	ProcessMsgResult ProcessMsgResult `xml:"ProcessMsgResult"`
}

type ErrorResponse struct {
	XMLName     xml.Name `xml:"ERROR"`
	Code        int64    `xml:"CODE"`
	Description string   `xml:"MSGTXT"`
}
