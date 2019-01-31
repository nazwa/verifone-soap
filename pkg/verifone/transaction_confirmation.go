package verifone

import (
	"encoding/xml"

	"github.com/shopspring/decimal"
)

type VgConfirmationRequest struct {
	XMLName xml.Name `xml:"vgconfirmationrequest"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Ns      string   `xml:"xmlns,attr"`

	// Session identifier
	SessionGUID string `xml:"sessionguid"`
	// TransactionID from Transaction request
	TransactionId string `xml:"transactionid"`
	// AuthCode if transaction was authorised offline
	OfflineAuthCode string `xml:"offlineauthcode,omitempty"`
	// Additional value to add to total (e.g. service tip)
	Gratuity *decimal.Decimal `xml:"gratuity,omitempty"`
}

func (this Client) ConfirmTransaction(sessionGuid, transactionId, offlineAuthCode string, gratuity *decimal.Decimal) (response VgTransactionResponse, err error) {
	v := VgConfirmationRequest{
		Xsi:             Xsi,
		Xsd:             Xsd,
		Ns:              Ns,
		SessionGUID:     sessionGuid,
		TransactionId:   transactionId,
		OfflineAuthCode: offlineAuthCode,
		Gratuity:        gratuity,
	}

	err = this.Call(MsgTypeConfirmTransaction, v, &response)

	return
}
