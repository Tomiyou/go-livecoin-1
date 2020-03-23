package livecoin

import "github.com/shopspring/decimal"

type DepositAddress struct {
	Currency string `json:"currency"`
	UserID   int    `json:"userId"`
	UserName string `json:"userName"`
	Wallet   string `json:"wallet"`
}

type Withdrawal struct {
	Fault                interface{}     `json:"fault"`
	UserId               int64           `json:"userId"`
	UserName             string          `json:"userName"`
	Id                   int64           `json:"id"`
	State                string          `json:"state"`
	CreateDate           int64           `json:"createDate"`
	LastModifyDate       int64           `json:"lastModifyDate"`
	VerificationType     string          `json:"verificationType"`
	VerificationData     interface{}     `json:"verificationData"`
	Comment              interface{}     `json:"comment"`
	Description          string          `json:"description"`
	Amount               decimal.Decimal `json:"amount"`
	Currency             string          `json:"currency"`
	AccountTo            string          `json:"accountTo"`
	AcceptDate           interface{}     `json:"acceptDate"`
	ValueDate            interface{}     `json:"valueDate"`
	DocDate              int64           `json:"docDate"`
	DocNumber            int64           `json:"docNumber"`
	CorrespondentDetails interface{}     `json:"correspondentDetails"`
	AccountFrom          string          `json:"accountFrom"`
	Outcome              bool            `json:"outcome"`
	External             interface{}     `json:"external"`
	ExternalKey          string          `json:"externalKey"`
	ExternalSystemId     int             `json:"externalSystemId"`
	ExternalServiceId    interface{}     `json:"externalServiceId"`
	Wallet               string          `json:"wallet"`
}
