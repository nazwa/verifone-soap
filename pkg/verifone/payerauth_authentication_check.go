package verifone

import (
	"encoding/xml"
)

type EnrollmentStatus string
type AuthenticationStatus string

const (
	EnrollmentStatusYes EnrollmentStatus = "Y"
	EnrollmentStatusNo  EnrollmentStatus = "N"

	AuthenticationStatusSuccess              AuthenticationStatus = "Y"
	AuthenticationStatusFailure              AuthenticationStatus = "N"
	AuthenticationStatusAttemptedNotEnrolled AuthenticationStatus = "A"
	AuthenticationStatusTechnicalProblem     AuthenticationStatus = "U"
)

type VgPayerAuthAuthenticationCheckRequest struct {
	XMLName xml.Name `xml:"vgpayerauthauthenticationcheckrequest"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Ns      string   `xml:"xmlns,attr"`

	// Session identifier
	SessionGUID string `xml:"sessionguid"`
	// Merchant can add a reference to cross reference responses relating to the same transaction
	MerchantReference string `xml:"merchantreference,omitempty"`
	// Unique Identifier returned during EnrollmentCheckResponse.
	PayerAuthRequestId int64 `xml:"payerauthrequestid"`
	// Compressed and encoded Payer Authentication Response message, returned in response from Visa / Mastercard (Only included if received)
	Pares string `xml:"pares"`
	// Indicates if the card was enrolled – Y/N
	Enrolled EnrollmentStatus `xml:"enrolled"`
}

type VgPayerAuthAuthenticationCheckResponse struct {
	ErrorCode        int64  `xml:"errorcode"`
	ErrorDescription string `xml:"errordescription"`

	// Session identifier
	SessionGUID string `xml:"sessionguid"`
	// Merchant can add a reference to cross reference responses relating to the same transaction
	MerchantReference string `xml:"merchantreference"`
	//Additional transaction security data
	AtsData string `xml:"atsdata"`
	// This property indicates whether the transaction has been authenticated or not. • Y – The customer was successfully authenticated. All data needed for clearing is included. • N – The customer failed authentication • A – Attempted processing. The APACS message will show verified enrolment but cardholder is not participating at this time. • U – Authentication could not be performed due to technical or other problems
	AuthenticationStatus AuthenticationStatus `xml:"authenticationstatus"`
	//The certificate that signed the Payer Authentication Response (PARes) message
	AuthenticationCertificate string `xml:"authenticationcertificate"`
	// This property contains a 28-byte Base-64 encoded Cardholder Authentication Verification Value (CAVV)
	AuthenticationCavv string `xml:"authenticationcavv"`
	// Two digit Electronic Commerce Indicator (ECI) value
	AuthenticationEci string `xml:"authenticationeci"`
	// The date and time in which the Payer Authentication Response (PARes) message was signed by the Access Control Server (ACS). The value is expressed in GMT and uses the format "YYYYMMDD HH:MM:SS"
	AuthenticationTime string `xml:"authenticationtime"`
}

func (this Client) PayerAuthAuthenticationCheck(sessionGuid, merchantReference string, payerAuthRequestId int64, pares string, enrollmentStatus EnrollmentStatus) (response VgPayerAuthAuthenticationCheckResponse, err error) {
	v := VgPayerAuthAuthenticationCheckRequest{
		Xsi:                Xsi,
		Xsd:                Xsd,
		Ns:                 Ns,
		SessionGUID:        sessionGuid,
		MerchantReference:  merchantReference,
		PayerAuthRequestId: payerAuthRequestId,
		Pares:              pares,
		Enrolled:           enrollmentStatus,
	}

	err = this.Call(MsgTypePayerAuthAuthenticationCheck, v, &response, nil)

	return
}
