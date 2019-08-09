package livecoin

import (
	"github.com/shopspring/decimal"
)

type Balance struct {
	Type     string          `json:"type"`
	Currency string          `json:"currency"`
	Value    decimal.Decimal `json:"value"`
}
