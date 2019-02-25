package verifone

import (
	"encoding/xml"

	"github.com/nazwa/verifone-soap/pkg/gosoap"
	"github.com/pkg/errors"
)

const (
	MsgTypeGenSession                   = "VGGENERATESESSIONREQUEST"
	MsgTypeRegisterToken                = "VGTOKENREGISTRATIONREQUEST"
	MsgTypeGetCardDetails               = "VGGETCARDDETAILSREQUEST"
	MsgTypeTransaction                  = "VGTRANSACTIONREQUEST"
	MsgTypeConfirmTransaction           = "VGCONFIRMATIONREQUEST"
	MsgTypeRejectTransaction            = "VGREJECTIONREQUEST"
	MsgTypePayerAuthEnrollmentCheck     = "VGPAYERAUTHENROLLMENTCHECKREQUEST"
	MsgTypePayerAuthAuthenticationCheck = "VGPAYERAUTHAUTHENTICATIONCHECKREQUEST"

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
	config       Config
	soap         *gosoap.Client
	processingDb string
}

func NewClient(cfg Config, processingDb string) *Client {
	if cfg.Url == "" {
		cfg.Url = "https://txn-cst.cxmlpg.com/XML4/commideagateway.asmx?WSDL"
	}

	soap, err := gosoap.SoapClient(cfg.Url)
	if err != nil {
		panic(err)
	}

	return &Client{
		config:       cfg,
		soap:         soap,
		processingDb: processingDb,
	}
}

func (this Client) getClientHeader() MessageClientHeader {
	return MessageClientHeader{
		SystemGUID:    this.config.SystemGUID,
		SystemID:      this.config.SystemID,
		Passcode:      this.config.Passcode,
		ProcessingDB:  this.processingDb,
		CDATAWrapping: false,
	}
}

type MessageClientHeader struct {
	SystemGUID    string `xml:"SystemGUID"`
	SystemID      string `xml:"SystemID"`
	Passcode      string `xml:"Passcode"`
	ProcessingDB  string `xml:"ProcessingDB"`
	CDATAWrapping bool   `xml:"CDATAWrapping"`
}

func (this Client) Call(msgType string, msgData interface{}, target interface{}, headerTarget *ClientHeaderResponse) (err error) {
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

	if headerTarget != nil {
		*headerTarget = *response.ProcessMsgResult.ClientHeader
	}

	return
}
