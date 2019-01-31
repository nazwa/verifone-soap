package verifone

import (
	"encoding/xml"
)

type VgGenerateSessionRequest struct {
	XMLName     xml.Name `xml:"vggeneratesessionrequest"`
	Xsi         string   `xml:"xmlns:xsi,attr"`
	Xsd         string   `xml:"xmlns:xsd,attr"`
	Ns          string   `xml:"xmlns,attr"`
	ReturnUrl   string   `xml:"returnurl"`
	FullCapture bool     `xml:"fullcapture"`
}

type VgGenerateSessionResponse struct {
	SessionGUID      string `xml:"sessionguid"`
	SessionPasscode  string `xml:"sessionpasscode"`
	ErrorCode        int64  `xml:"errorcode"`
	ErrorDescription string `xml:"errordescription"`
}

func (this Client) BeginSession() (response VgGenerateSessionResponse, err error) {

	v := VgGenerateSessionRequest{
		Xsi:         "http://www.w3.org/2001/XMLSchema-instance",
		Xsd:         "http://www.w3.org/2001/XMLSchema",
		Ns:          "VANGUARD",
		ReturnUrl:   "http://google.com",
		FullCapture: true,
	}

	err = this.call(MsgTypeGenSession, v, &response)

	return
}
