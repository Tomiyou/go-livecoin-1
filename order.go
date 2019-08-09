package livecoin

import (
	"github.com/shopspring/decimal"
)

type NewOrder struct {
	Success bool   `json:"success"`
	Added   bool   `json:"added"`
	OrderId uint64 `json:"orderId"`
}

type CancelledOrder struct {
	Success       bool            `json:"success"`
	Cancelled     bool            `json:"cancelled"`
	Exception     string          `json:"exception"`
	Quantity      decimal.Decimal `json:"quantity"`
	TradeQuantity decimal.Decimal `json:"tradeQuantity"`
}

type OrderInfo struct {
	Id                uint64          `json:"id"`
	ClientId          uint64          `json:"client_id"`
	Status            string          `json:"status"`
	TickerPair        string          `json:"symbol"`
	Price             decimal.Decimal `json:"price"`
	Quantity          decimal.Decimal `json:"quantity"`
	RemainingQuantity decimal.Decimal `json:"remaining_quantity"`
	Blocked           decimal.Decimal `json:"blocked"`
	BlockedRemain     decimal.Decimal `json:"blocked_remain"`
	CommissionRate    decimal.Decimal `json:"commission_rate"`
	Trades            interface{}     `json:"trades"`
}
