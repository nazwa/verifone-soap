package verifone

import (
	"encoding/xml"
)

type VgRejectionRequest struct {
	XMLName xml.Name `xml:"vgrejectionrequest"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Ns      string   `xml:"xmlns,attr"`

	// Session identifier
	SessionGUID string `xml:"sessionguid"`
	// TransactionID from Transaction request
	TransactionId string `xml:"transactionid"`
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
	//Amex Card – 3 or 4 digits (front of card)  All Other Cards – 3 or 4 digits (rear security strip)
	Csc string `xml:"csc,omitempty"`
	// Field checked by Address Verification System (AVS) add on module, ignored if module not enabled. AVS configuration can make this field mandatory. Numerics from house name\number
	AvsHouse string `xml:"avshouse,omitempty"`
	// Field checked by Address Verification System (AVS) add on module, ignored if module not enabled. AVS configuration can make this field mandatory. Numerics from postcode only
	AvsPostCode string `xml:"avspostcode,omitempty"`
}

func (this Client) RejectTransaction(sessionGuid, transactionId string) (response VgTransactionResponse, err error) {
	v := VgRejectionRequest{
		Xsi:           Xsi,
		Xsd:           Xsd,
		Ns:            Ns,
		SessionGUID:   sessionGuid,
		TransactionId: transactionId,
	}

	err = this.Call(MsgTypeRejectTransaction, v, &response, nil)

	return
}
