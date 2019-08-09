package livecoin

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"time"
)

type Trade struct {
	Id            uint64          `json:"id"`
	ClientOrderId uint64          `json:"clientorderid"`
	Type          string          `json:"type"`
	Symbol        string          `json:"symbol"`
	Date          time.Time       `json:"datetime"`
	Price         decimal.Decimal `json:"price"`
	Quantity      decimal.Decimal `json:"quantity"`
	Commission    decimal.Decimal `json:"commission"`
}

func (t *Trade) UnmarshalJSON(data []byte) error {
	var err error
	type Alias Trade
	aux := &struct {
		Date int64 `json:"datetime"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t.Date = time.Unix(aux.Date, 0)
	return nil
}
