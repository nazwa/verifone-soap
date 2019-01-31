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

type ProcessMsgResult struct {
	MsgType string `xml:"MsgType"`
	MsgData []byte `xml:"MsgData"`
}

type ProcessMsgResponse struct {
	ProcessMsgResult ProcessMsgResult `xml:"ProcessMsgResult"`
}

type ErrorResponse struct {
	XMLName     xml.Name `xml:"ERROR"`
	Code        int64    `xml:"CODE"`
	Description string   `xml:"MSGTXT"`
}
