package verifone

import (
	"encoding/xml"
)

type CpcOption int64

const (
	CpcOptionNotCpc       CpcOption = 0
	CpcOptionCpc          CpcOption = 1
	CpcOptionMaskRequired CpcOption = 2
)

type VgGetCardDetailsRequest struct {
	XMLName xml.Name `xml:"vggetcarddetailsrequest"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Ns      string   `xml:"xmlns,attr"`

	// Session identifier
	SessionGUID string `xml:"sessionguid"`
}

type VgGetCardDetailsResponse struct {
	ErrorCode        int64  `xml:"errorcode"`
	ErrorDescription string `xml:"errordescription"`

	// Session identifier
	SessionGUID string `xml:"sessionguid"`
	// Echo from request
	FullCapture bool `xml:"fullcapture"`
	// Card scheme identifier
	MkCardSchemeId int64 `xml:"mkcardschemeid"`
	// Card scheme name
	SchemeName string `xml:"schemename"`
	// Length of issue number, if required
	IssueNoLength int64 `xml:"issuenolength"`
	// Indicates if start date required
	StartDateRequired bool `xml:"startdaterequired"`
	// Length of Card Security Code, if required
	CscLength string `xml:"csclength"`
	//Indicates if Payer Authentication is supported
	AllowPayerAuth bool `xml:"allowpayerauth"`
	// Available Values: 0 = Not CPC Card 1 = CPC Card 2 = Mask Identification Required
	CpcOption CpcOption `xml:"cpcoption"`
	// Mask used to identify if the PAN matches the format of a CPC card or not.  Valid characters are:
	// _ = A single numeric digit 0-9 = Specific numeric digit % = Any combination of numeric digits
	CpcIndentificationMask string `xml:"cpcidentificationmask"`
	// Starred PAN
	PanStar string `xml:"panstar"`
	// SHA-256 hashed representation of the PAN
	CardNumberHash string `xml:"cardnumberhash"`
	// Expiry date from customer’s post Only returned if capture method was ‘full capture’
	ExpiryDate string `xml:"expirydate"`
	// Start date from customer’s post Only returned if capture method was ‘full capture’
	StartDate string `xml:"startdate"`
	// Issue no from customer’s post Only returned if capture method was ‘full capture’
	IssueNo string `xml:"issueno"`
	// Cardholder’s billing address first line from customer’s post
	Address1 string `xml:"address1"`
	// Cardholder’s billing address second line from customer’s post
	Address2 string `xml:"address2"`
	// Cardholder’s billing address town from customer’s post
	Town string `xml:"town"`
	// Cardholder’s billing address county from customer’s post
	County string `xml:"county"`
	// Cardholder’s billing address postcode from customer’s post
	PostCode string `xml:"postcode"`
	// Cardholder’s billing address country from customer’s post
	Country string `xml:"country"`
	// Cardholder’s name
	CardholderName string `xml:"cardholdername"`
}

func (this Client) GetCardDetails(sessionGuid string) (response VgGetCardDetailsResponse, err error) {
	v := VgGetCardDetailsRequest{
		Xsi:         Xsi,
		Xsd:         Xsd,
		Ns:          Ns,
		SessionGUID: sessionGuid,
	}

	err = this.Call(MsgTypeGetCardDetails, v, &response)

	return
}
