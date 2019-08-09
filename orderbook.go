package livecoin

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
)

type Orderbook struct {
	Asks []OrderbookPair `json:"asks"`
	Bids []OrderbookPair `json:"bids"`
	Timestamp int64      `json:"timestamp"`
}

type OrderbookPair struct {
	Price     decimal.Decimal
	Quantity  decimal.Decimal
	Timestamp int64
}

func (p *OrderbookPair) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&p.Price, &p.Quantity, &p.Timestamp}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), 3; g != e {
		return fmt.Errorf("wrong number of fields in OrderbookPair: %d != %d", g, e)
	}
	return nil
}
