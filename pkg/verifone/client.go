package verifone

import (
	"encoding/xml"
	"verifone-soap/pkg/gosoap"

	"github.com/pkg/errors"
)

const (
	MsgTypeGenSession    = "VGGENERATESESSIONREQUEST"
	MsgtypeRegisterToken = "VGTOKENREGISTRATIONREQUEST"

	Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	Xsd = "http://www.w3.org/2001/XMLSchema"
	Ns  = "VANGUARD"
)

type Config struct {
	SystemID   string
	SystemGUID string
	Passcode   string
	Url        string
}

type Client struct {
	config Config
	soap   *gosoap.Client
}

func NewClient(cfg Config) *Client {
	if cfg.Url == "" {
		cfg.Url = "https://txn-cst.cxmlpg.com/XML4/commideagateway.asmx?WSDL"
	}

	soap, err := gosoap.SoapClient(cfg.Url)
	if err != nil {
		panic(err)
	}

	return &Client{
		config: cfg,
		soap:   soap,
	}
}

func (this Client) getClientHeader() MessageClientHeader {
	return MessageClientHeader{
		SystemGUID: this.config.SystemGUID,
		SystemID:   this.config.SystemID,
		Passcode:   this.config.Passcode,
	}
}

type MessageClientHeader struct {
	SystemGUID string `xml:"SystemGUID"`
	SystemID   string `xml:"SystemID"`
	Passcode   string `xml:"Passcode"`
}

func (this Client) call(msgType string, msgData interface{}, target interface{}) (err error) {
	var body []byte
	body, err = xml.Marshal(msgData)
	if err != nil {
		return
	}

	// Verifone uses a single ProcessMsg method, so this is a simplified wrapper around it
	msg := ProcessMsg{
		Ns: "https://www.commidea.webservices.com",
		Message: Message{
			ClientHeader: this.getClientHeader(),
			MsgType:      msgType,
			MsgData:      MsgData{Content: body},
		},
	}

	if err = this.soap.Call(msg); err != nil {
		return
	}

	response := ProcessMsgResponse{}
	if err = this.soap.Unmarshal(&response); err != nil {
		return
	}

	if response.ProcessMsgResult.MsgType == "ERROR" {
		errResp := ErrorResponse{}
		if err = xml.Unmarshal(response.ProcessMsgResult.MsgData, &errResp); err != nil {
			// This is a super bad place
			// We should never land here
			panic(errors.Wrap(err, "This should never happen!!!!"))
		}

		return errors.Errorf("[%d]: %s", errResp.Code, errResp.Description)
	}

	err = xml.Unmarshal(response.ProcessMsgResult.MsgData, &target)
	return
}

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
