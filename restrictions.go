package livecoin

import (
	"github.com/shopspring/decimal"
)

type Restrictions struct {
	Success      bool            `json:"success"`
	MinBtcVolume decimal.Decimal `json:"minBtcVolume"`
	Restrictions []Restriction   `json:"restrictions"`
}

type Restriction struct {
	CurrencyPair     string          `json:"currencyPair"`
	MinLimitQuantity decimal.Decimal `json:"minLimitQuantity"`
	PriceScale       int             `json:"priceScale"`
}
