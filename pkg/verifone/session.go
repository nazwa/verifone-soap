package verifone

import (
	"encoding/xml"
)

type VgGenerateSessionRequest struct {
	XMLName xml.Name `xml:"vggeneratesessionrequest"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Ns      string   `xml:"xmlns,attr"`

	// URL to redirect the customers browser to after receiving the post
	ReturnUrl string `xml:"returnurl"`
	// Specifies if the customer will be posting all the required card details or just partial card details
	// Note: if only capturing partial card details, it will be necessary for the merchant to collect the remaining details and supply within the transaction request
	FullCapture bool `xml:"fullcapture"`
}

type VgGenerateSessionResponse struct {
	ErrorCode        int64  `xml:"errorcode"`
	ErrorDescription string `xml:"errordescription"`

	// Session identifier. 32 character hexadecimal string
	SessionGUID string `xml:"sessionguid"`
	// Session passcode. 32 character hexadecimal string
	SessionPasscode string `xml:"sessionpasscode"`

	// Manually added processingDB
	ProcessingDB string
}

func (this Client) BeginSession(returnUrl string, fullCapture bool) (response VgGenerateSessionResponse, err error) {
	v := VgGenerateSessionRequest{
		Xsi:         Xsi,
		Xsd:         Xsd,
		Ns:          Ns,
		ReturnUrl:   returnUrl,
		FullCapture: fullCapture,
	}

	clientHeader := &ClientHeaderResponse{}

	err = this.Call(MsgTypeGenSession, v, &response, clientHeader)
	if err == nil {
		response.ProcessingDB = clientHeader.ProcessingDB
	}
	return
}
