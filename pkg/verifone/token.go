package verifone

import (
	"encoding/xml"
)

type VgTokenRegistrationRequest struct {
	XMLName xml.Name `xml:"vgtokenregistrationrequest"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Ns      string   `xml:"xmlns,attr"`

	// Session identifier
	SessionGUID string `xml:"sessionguid"`
	// Merchant can add a reference to cross reference responses relating to the same transaction
	MerchantReference string `xml:"merchantreference,omitempty"`
	// Card expiry month and year (YYMM)
	ExpiryDate string `xml:"expirydate,omitempty"`
	// Allow purchase transaction type
	Purchase bool `xml:"purchase"`
	// Allow refund transaction type
	Refund bool `xml:"refund"`
	// Allow cashback transaction type
	CashBack bool `xml:"cashback"`
	// Last date on which the token can be utilised.  Format of date to be:  DDMMCCYY
	TokenExpirationDate string `xml:"tokenexpirationdate"`
}

type VgTokenRegistrationResponse struct {
	ErrorCode        int64  `xml:"errorcode"`
	ErrorDescription string `xml:"errordescription"`

	// Session identifier
	SessionGUID string `xml:"sessionguid"`
	// Merchant can add a reference to cross reference responses relating to the same transaction
	MerchantReference string `xml:"merchantreference"`
	// Unique identifier for registered PAN. The maximum size limit for this field is 18.
	// This value should only be stored when the error code field contains a ‘0’. For all other error conditions, the tokenID should not be stored.
	TokenId string `xml:"tokenid"`
	// Name of card scheme
	CardSchemeName string `xml:"cardschemename"`
}

func (this Client) RegisterToken(sessionGuid, merchantReference, expiryDate, tokenExpirationDate string, purchase, refund, cashback bool) (response VgTokenRegistrationResponse, err error) {
	v := VgTokenRegistrationRequest{
		Xsi:                 Xsi,
		Xsd:                 Xsd,
		Ns:                  Ns,
		SessionGUID:         sessionGuid,
		MerchantReference:   merchantReference,
		ExpiryDate:          expiryDate,
		TokenExpirationDate: tokenExpirationDate,
		Purchase:            purchase,
		Refund:              refund,
		CashBack:            cashback,
	}

	err = this.Call(MsgTypeRegisterToken, v, &response)

	return
}
