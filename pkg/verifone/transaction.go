package verifone

import (
	"encoding/xml"

	"github.com/shopspring/decimal"
)

type TxnType string
type ApacsCapability string
type CaptureMethod int64
type ProcessingIdentifier int64
type SchemeID int64
type TxnResult string
type AvsResult int64
type CvcResult int64

const (
	TxnTypePurchase             TxnType = "01"
	TxnTypeRefund               TxnType = "02"
	TxnTypeCashAdvance          TxnType = "04"
	TxnTypePurchaseWithCashBack TxnType = "05"
	TxnTypeContinousAuthority   TxnType = "06"
	TxnTypeAccountCheck         TxnType = "07"

	ApacsCapabilitySwipedUnattended                       ApacsCapability = "3291"
	ApacsCapabilityMailOrTelephoneOrder                   ApacsCapability = "4290"
	ApacsCapabilityCnpOrEcommerce                         ApacsCapability = "4298"
	ApacsCapabilityKeyedAndSwipedAttended                 ApacsCapability = "6290"
	ApacsCapabilityContactIccKeyedAndSwiped               ApacsCapability = "7296"
	ApacsCapabilitySwipedOrIccOrContactlessUnattended     ApacsCapability = "B291"
	ApacsCapabilityContactlessAndKeyed                    ApacsCapability = "C296"
	ApacsCapabilityKeyedOrSwipedOrContactOrContactlessEmv ApacsCapability = "F296"

	CaptureMethodKeyedAttended              CaptureMethod = 1
	CaptureMethodKeyedUnattendedOrMailOrder CaptureMethod = 2
	CaptureMethodSwiped                     CaptureMethod = 3
	CaptureMethodIccOrSwiped                CaptureMethod = 4
	CaptureMethodIccOrSignature             CaptureMethod = 5
	CaptureMethodIccPinOnly                 CaptureMethod = 6
	CaptureMethodIccPinAndSignature         CaptureMethod = 7
	CaptureMethodIccNoCvm                   CaptureMethod = 8
	CaptureMethodContactlessEmv             CaptureMethod = 9
	CaptureMethodContactlessMagStripe       CaptureMethod = 10
	CaptureMethodPhoneOrderUnattended       CaptureMethod = 11
	CaptureMethodEcommerceUnattended        CaptureMethod = 12

	ProcessingIdentifierAuthAndCharge ProcessingIdentifier = 1
	ProcessingIdentifierAuthOnly      ProcessingIdentifier = 2
	ProcessingIdentifierChargeOnly    ProcessingIdentifier = 3

	SchemeIDAmex                                SchemeID = 1
	SchemeIDVisaOrRoiVisaDebit                  SchemeID = 2
	SchemeIDMasterCardOrMasterCardOne           SchemeID = 3
	SchemeIDMaestro                             SchemeID = 4
	SchemeIDVisaDebitOrRoiVisaDebitOrVisaCredit SchemeID = 6
	SchemeIDAmexCpc                             SchemeID = 36
	SchemeIDMasterCardDebit                     SchemeID = 49
	SchemeIDInvalid                             SchemeID = 999

	TxnResultError       TxnResult = "ERROR"
	TxnResultReferral    TxnResult = "REFERRAL"
	TxnResultCommsDown   TxnResult = "COMMSDOWN"
	TxnResultDeclined    TxnResult = "DECLINED"
	TxnResultRejected    TxnResult = "REJECTED"
	TxnResultCharged     TxnResult = "CHARGED"
	TxnResultApproved    TxnResult = "APPROVED"
	TxnResultAuthorised  TxnResult = "AUTHORISED"
	TxnResultAuthOnly    TxnResult = "AUTHONLY"
	TxnResultVerified    TxnResult = "VERIFIED"
	TxnResultNotVerifier TxnResult = "NOT VERIFIED"

	AvsResultNotProvided  AvsResult = 0
	AvsResultNotChecked   AvsResult = 1
	AvsResultMatched      AvsResult = 2
	AvsResultNotMatched   AvsResult = 4
	AvsResultPartialMatch AvsResult = 8

	CvcResultNotProvided CvcResult = 0
	CvcResultNotChecked  CvcResult = 1
	CvcResultMatched     CvcResult = 2
	CvcResultNotMatched  CvcResult = 4
)

type PayerAuthAuxiliaryData struct {
	XMLName xml.Name `xml:"payerauthauxiliarydata"`
	// Indicates if the transaction authenticated or not:
	// Y – Customer was successfully authenticated
	// N – Customer failed authentication, and the transaction declined
	// A – Attempts processing. APACS message will show verified enrollment but cardholder not participating
	// U – Enrollment could not be completed, due to technical or other problem authenticationcavv
	AuthenticationStatus AuthenticationStatus `xml:"authenticationstatus"`
	// Contains 28-byte Base-64 encoded Cardholder Authentication Verification Value (CAVV)
	AuthenticationCavv string `xml:"authenticationcavv"`
	// 2 digit Electronic Commerce Indicator (ECI) value
	AuthenticationEci string `xml:"authenticationeci"`
	// Data to populate authorisation message
	AtsData string `xml:"atsdata"`
	// TransactionID should be populated with the PayerAuthRequestID provided in the PayerAuth EnrollmentCheck Response
	TransactionID int64 `xml:"transactionid"`
}

type VgTransactionRequest struct {
	XMLName xml.Name `xml:"vgtransactionrequest"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Ns      string   `xml:"xmlns,attr"`

	// Session identifier
	SessionGUID string `xml:"sessionguid"`
	// Merchant can add a reference to cross reference responses relating to the same transaction
	MerchantReference string `xml:"merchantreference,omitempty"`
	// Account reference number, supplied by Verifone
	AccountID int64 `xml:"accountid"`
	// 01 – Purchase
	// 02 – Refund
	// 04 – Cash Advance
	// 05 – Purchase with cash back (PWCB)
	// 06 – Continuous Authority
	// 07 – Account Check (for more details, please refer to the section within this guide)
	TxnType TxnType `xml:"txntype"`
	// This is the three digit currency code (numeric).
	TransactionCurrencyCode string `xml:"transactioncurrencycode"`
	// In accordance with the numeric values defined in ISO 3166 (see Appendix C)
	TerminalCountryCode string `xml:"terminalcountrycode"`
	// This is the functionality supported by the terminal in the format of that defined by the APACS standard.
	// These are:
	// 3291 – Only Swiped and Contact ICC unattended
	// 4290 – Mail Order/Telephone Order
	// 4298 – CNP/ECommerce (if flagged for payer authorisation with acquirer; no CNP transactions are allowed with the exception of refunds)
	// 6290 - Keyed and Swiped Customer Present
	// 7296 – Contact (ICC) Keyed and Swiped
	// B291 – Swiped, Contact ICC and Contactless unattended
	// C296 – Contactless and keyed transactions (a contactless auxiliary record should be presented for all transactions passed under this terminal type)
	// F296 – Keyed, Swiped, Contact and Contactless EMV transactions (a contactless auxiliary record should be present for all transactions passed under this terminal type). Integrators should check with Technical Services to confirm that they have the correct capabilities.
	ApacsTerminalCapabilities ApacsCapability `xml:"apacsterminalcapabilities"`
	// This indicates how the card details were obtained. Acceptable values are:
	// 1 – Keyed Cardholder Present
	// 2 – Keyed Cardholder Not Present Mail Order
	// 3 – Swiped
	// 4 – ICC Fallback to Swipe
	// 5 – ICC Fallback to Signature
	// 6 – ICC PIN Only
	// 7 – ICC PIN and Signature
	// 8 – ICC – No CVM
	// 9 – Contactless EMV
	// 10 – Contactless Mag Stripe
	// 11 – Keyed Cardholder Not Present Telephone Order
	// 12 – Keyed Cardholder Not Present Ecommerce Order
	CaptureMethod CaptureMethod `xml:"capturemethod"`
	// This indicates the type of processing that needs to be undertaken. Current available values are as follows:
	// 1 – Auth and Charge
	// 2 – Auth Only
	// 3 – Charge Only
	// All refund transactions should use the ‘Charge Only’ option.
	ProcessingIdentifier ProcessingIdentifier `xml:"processingidentifier"`
	//Amex Card – 3 or 4 digits (front of card)  All Other Cards – 3 or 4 digits (rear security strip)
	Csc string `xml:"csc,omitempty"`
	// Field checked by Address Verification System (AVS) add on module, ignored if module not enabled. AVS configuration can make this field mandatory. Numerics from house name\number
	AvsHouse string `xml:"avshouse,omitempty"`
	// Field checked by Address Verification System (AVS) add on module, ignored if module not enabled. AVS configuration can make this field mandatory. Numerics from postcode only
	AvsPostCode int64 `xml:"avspostcode,omitempty"`
	// Card expiry month and year (YYMM)
	ExpiryDate string `xml:"expirydate"`
	// Total value of transaction including tax. Applies to: Purchase, Refund, Cheque Guarantee, Cash Advance, and Purchase with Cash Back. With PWCB, field should only contain the values of the goods or services provided. Decimal point recommended but optional, e.g.: 1.23 = £1.23 123 = £123 000001.23 = £1.23 Only positive values. Values will be truncated to the correct number of decimal places required for the transaction currency (set by the merchant account being used)
	// Please note: For Account Check transaction, the value must be in the following format: “0.00”.
	TxnValue decimal.Decimal `xml:"txnvalue"`
	// Total Cash Back value for PWCB transactions. Values will be truncated (without rounding) to the number of decimal places required for the transaction currency. Positive values only.
	CashbackValue decimal.Decimal `xml:"cashbackvalue,omitempty"`
	// Additional value to add to total (e.g. service tip)
	Gratuity decimal.Decimal `xml:"gratuity,omitempty"`
	// Only supplied for Offline transactions
	AuthCode string `xml:"authcode,omitempty"`
	// Date and time the transaction was started, based on GMT (DD/MM/YYYY HH:MM:SS).
	TransactionDateTime string `xml:"transactiondatetime,omitempty"`
	// In accordance with the numeric values defined in ISO 3166
	TransactionCountryCode string `xml:"transactioncountrycode"`
	// Employee identifier
	EmployeeId string `xml:"employeeid,omitempty"`
	// Payer Authentication auxiliary data. This field is conditional upon the capture method/transaction type. If Payer Authentication is performed this data must be supplied, even for non-supporting card schemes. Capture methods such as ICC will not require Payer Auth auxiliary data to be supplied
	PayerAuthAuxiliaryData *PayerAuthAuxiliaryData `xml:"payerauthauxiliarydata,omitempty"`
	// Denotes if the transaction is a procurement card/VGIS transaction.
	VgisTransaction *bool `xml:"vgistransaction"`
	// Account passcode, supplied by Verifone
	AccountPasscode string `xml:"accountpasscode"`
	// Specifies whether to return the ‘gdethash’ and ‘panstar’ fields within the transaction response message.
	// Valid values are: 0 = False 1 = True
	ReturnHash bool `xml:"returnhash"`
}

type VgTransactionResponse struct {
	ErrorCode        int64  `xml:"errorcode"`
	ErrorDescription string `xml:"errordescription"`

	// Session identifier
	SessionGUID string `xml:"sessionguid"`
	// Merchant can add a reference to cross reference responses relating to the same transaction
	MerchantReference string `xml:"merchantreference"`
	// Transaction ID (unique only to the processing database utilised)
	TransactionId string `xml:"transactionid"`
	// Extended date and time string ( YYYY-MM-DDTHH:MM:SS:ss)
	ResultDateTimeString string `xml:"resultdatetimestring"`
	// Unique merchant number
	MerchantNumber string `xml:"merchantnumber"`
	// Terminal ID
	Tid string `xml:"tid"`
	// Card scheme name 1 – Amex 2 – Visa / ROI Visa Debit 3 – MasterCard / MasterCard One 4 – Maestro 5 – Diners / Discover 6 – Visa Debit / ROI Visa Debit / Visa Credit 7 – JCB 8 – BT Test Host 9 – Time / TradeUK Account card 10 – Solo (ceased) 11 – Electron 21 – Visa CPC 23 – AllStar CPC 24 – EDC/Maestro (INT) 25 – Laser 26 – LTF 27 – CAF (Charity Aids Foundation) 28 – Creation (Sears / Duet) 29 – Clydesdale Financial Services 30 – Style card 31 – BHS Gold 32 – Mothercare Card 33 – Arcadia Group cards (Privilege, Shareholder & Staff) 35 – BA AirPlus 36 – Amex CPC 41 – FCUK card (Style) 48 – Premier Inn Business Account card 49 – MasterCard Debit 50 – Stax Charge card 51 – IKEA Home card (IKANO) 52 – MasterCard One 53 – HFC Store card 999 – Invalid Card Range
	SchemeID SchemeID `xml:"SchemeID"`
	// Transaction message number (equivalent of EFTSN from previous versions of the Web Service)
	MessageNumber string `xml:"messagenumber"`
	// The Authorisation code that is returned by the bank. This will be blank if the transaction is declined and if the transaction value is below the floor limit.
	AuthCode string `xml:"authcode"`
	// The authmessage wil depend on the txnresult received, e.g. if ‘AuthOnly’ is received then the authmessage will return “ACCOUNT VALID” for Account Check transaction.
	AuthMessage string `xml:"authmessage"`
	// Telephone number to be called by the operator to seek manual authorisation. Only supplied for referred transactions
	VrTel string `xml:"vrtel"`
	// Transaction result: ERROR REFERRAL COMMSDOWN DECLINED REJECTED CHARGED APPROVED AUTHORISED AUTHONLY VERIFIED NOT VERIFIE
	TxnResult TxnResult `xml:"txnresult"`
	// Postcode AVS result: 0 – Not provided* 1 – Not checked 2 – Matched 4 – Not matched 8 – Partial Match *Default result when no details are provided
	PcAvsResult AvsResult `xml:"pcavsresult"`
	// Address line 1 AVS result: 0 – Not provided* 1 – Not checked 2 – Matched 4 – Not matched 8 – Partial Match *Default result when no details are provided
	AdAvsResult AvsResult `xml:"ad1avsresult"`
	// CVC result: 0 – Not provided* 1 – Not checked 2 – Matched 4 – Not matched *Default result when no details are provided
	CvcResult CvcResult `xml:"cvcresult"`
	// Acquirer response code. Should integrators wish to utilise this information to provide further insight into the transaction result, please contact your acquirer for further information, as this differs per acquirer.
	Arc string `xml:"arc"`
	// This indicates who actually performed the authorisation processing. Valid values are as follows: Not Provided = 0 Merchant = 1 Acquirer = 2 Card Scheme = 4 Issuer = 8
	AuthorisingEntity string `xml:"authorisingentity"`
	// VGIS Reference assigned to the transaction. Only returned when vgistransaction is passed within the transaction request as ‘true’
	VgisReference string `xml:"vgisreference"`
	// Hashed version of the card number, specific to the configuration of the merchant in question. Feature must be enabled before the field will be returned
	CustomerSpecifiedHash string `xml:"customerspecifiedhash"`
	// Starred version of the PAN.  Only returned if ‘returnhash’ set to true within the transaction request message
	PanStar string `xml:"panstar"`
	// SHA-256 representation of the PAN provided within the transaction request.  Only returned if ‘returnhash’ set to true within the transaction request message
	GdetHash string `xml:"gdethash"`
}

func (this Client) TransactionRequest(v VgTransactionRequest) (response VgTransactionResponse, err error) {
	v.Xsi = Xsi
	v.Xsd = Xsd
	v.Ns = Ns

	err = this.Call(MsgTypeTransaction, v, &response, nil)

	return
}
