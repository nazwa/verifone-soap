package verifone

import (
	"encoding/xml"
)

type Acquirer int64

const (
	AcquirerBarclaycardBusinessSterling          Acquirer = 1
	AcquirerNatWestStreamline                    Acquirer = 2
	AcquirerHMS                                  Acquirer = 3
	AcquirerLloydsTsbCardnet                     Acquirer = 4
	AcquirerElavon                               Acquirer = 5
	AcquirerBankOfScotland                       Acquirer = 6
	AcquirerAmericanExpress                      Acquirer = 7
	AcquirerClydesdaleBank                       Acquirer = 8
	AcquirerBarclaycardBusinessMultiCurrency     Acquirer = 9
	AcquirerBankOfIreland                        Acquirer = 10
	AcquirerNorthernBank                         Acquirer = 11
	AcquirerYorkshireBank                        Acquirer = 12
	AcquirerGECapital                            Acquirer = 13
	AcquirerUlsterBank                           Acquirer = 14
	AcquirerIntlBarclaycardBusinessSterling      Acquirer = 15
	AcquirerIntlLloydsTsbCardnet                 Acquirer = 16
	AcquirerIntlHms                              Acquirer = 17
	AcquirerIntlNatWest                          Acquirer = 18
	AcquirerIntlBArclaycardBusinessMultiCurrency Acquirer = 19
	AcquirerDiners                               Acquirer = 20
	AcquirerCreation                             Acquirer = 21
	AcquirerJCB                                  Acquirer = 23
	AcquirerAIB                                  Acquirer = 24
)

type VgPayerAuthEnrollmentCheckRequest struct {
	XMLName xml.Name `xml:"vgpayerauthenrollmentcheckrequest "`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Ns      string   `xml:"xmlns,attr"`

	// Session identifier
	SessionGUID string `xml:"sessionguid"`
	// Merchant can add a reference to cross reference responses relating to the same transaction
	MerchantReference string `xml:"merchantreference,omitempty"`
	// Account reference number, supplied by Verifone
	MkAccountId int64 `xml:"mkaccountid"`
	// Acquirer reference number 1 – Barclaycard Business (BMS) [Sterling only] 2 – NatWest Streamline 3 – HMS (HSBC) 4 – Lloyds TSB Cardnet 5 – Elavon (GiroBank) 6 – Bank Of Scotland 7 – American Express 8 – Clydesdale Bank 9 – Barclaycard Business (BMS) MultiCurrency 10 – Bank of Ireland 11 – Northern Bank 12 – Yorkshire Bank 13 – GE Capital 14 – Ulster Bank 15 – Int'l Barclaycard Business (BMS) [Sterling] 16 – Int'l Lloyds TSB Cardnet 17 – Int'l HMS (HSBC) 18 – Int'l NatWest 19 – Int'l Barclaycard Business (BMS) Multi 20 – Diners 21 – Creation 23 – JCB 24 – AIB
	MkAcquirerId Acquirer `xml:"mkacquirerid"`
	//The MerchantName must match the name shown online to the cardholder at the merchant’s site and the name submitted by the merchant’s acquirer in the settlement transaction
	MerchantName string `xml:"merchantname"`
	//This field contains a three digit number assigned by the signing member or processor to identify the merchant’s location country. Based on ISO Country Codes – 3166.
	MerchantCountryCode string `xml:"merchantcountrycode"`
	//This field contains the fully qualified URL of the  merchant site
	MerchantUrl string `xml:"merchanturl"`
	//This field contains a six digit assigned Bank Identification Number issued by the merchant’s member bank or processor. The acquirer Bank Identification Number (BIN) identifies the member bank that signed the merchant using the Point of Sale application
	VisaMerchantBankId string `xml:"visamerchantbankid,omitempty"`
	// This field contains a unique ID number which is assigned by the signing merchant’s acquirer, bank or processor. This field is used to identify the merchant within the VisaNet system
	VisaMerchantNumber string `xml:"visamerchantnumber,omitempty"`
	// This field contains a six digit assigned Bank Identification Number issued by the merchant’s member bank or processor. The acquirer Bank Identification Number (BIN) identifies the member bank that signed the merchant using the Point of Sale application
	McmMerchantBankId string `xml:"mcmmerchantbankid,omitempty"`
	//This field contains a unique ID number which is assigned by the signing merchant’s acquirer, bank or processor. This field is used to identify the merchant within the SecureCode system
	McmMerchantNumver string `xml:"mcmmerchantnumber,omitempty"`
	// Expiry date. Mandatory for partial capture but not required for full capture
	ExpiryDate string `xml:"expirydate,omitempty"`
	// This field contains a three digit number assigned by the signing member or processor to identify the merchant's authorisation currency. Based on ISO Country Code – 3166
	CurrencyCode string `xml:"currencycode"`
	// No of decimal places in currency field ie. GBP will be 2
	CurrencyExponent string `xml:"currencyexponent"`
	// This field contains the exact content of the HTTP accept header as sent to the merchant from the cardholder's user agent. This field is required only if the cardholder's user agent supplied a value.
	BrowserAcceptHeader string `xml:"browseracceptheader"`
	// This field contains the exact content of the HTTP useragent header as sent to the merchant from the cardholder's user agent. This field is only required if the cardholder's user agent supplied a value.
	BrowserUserAgentHeader string `xml:"browseruseragentheader"`
	// Amount to be authorised with implied decimal point ie. £10.00 is represented as 1000 and 0.10 is represented as 10.
	TransactionAmount string `xml:"transactionamount"`
	//The transaction amount is to be presented with all currencyspecific punctuation as this will be the number displayed to the customer. E.g. 10.00
	TransactionDisplayAmount string `xml:"transactiondisplayamount"`
	// This field contains a description of the goods or services being purchased, determined by the merchant
	TransactionDescription string `xml:"transactiondescription,omitempty"`
}

type VgPayerAuthEnrollmentCheckResponse struct {
	ErrorCode        int64  `xml:"errorcode"`
	ErrorDescription string `xml:"errordescription"`

	// Session identifier
	SessionGUID string `xml:"sessionguid"`
	// Merchant can add a reference to cross reference responses relating to the same transaction
	MerchantReference string `xml:"merchantreference"`
	//Unique Identifier
	PayerAuthRequestId int64 `xml:"payerauthrequestid"`
	// Indicates if card is enrolled in the 3D secure program.
	Enrolled EnrollmentStatus `xml:"enrolled"`
	// Fully qualified URL of an Access Control Server.
	AcsUrl string `xml:"acsurl"`
	// This field will contain the entire XML response packet from the Directory Server.
	Pareq string `xml:"pareq"`
}

func (this Client) PayerAuthEnrollmentCheck(v VgPayerAuthEnrollmentCheckRequest) (response VgPayerAuthEnrollmentCheckResponse, err error) {
	v.Xsi = Xsi
	v.Xsd = Xsd
	v.Ns = Ns

	err = this.Call(MsgTypePayerAuthEnrollmentCheck, v, &response)

	return
}
