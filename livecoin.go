// Package Livecoin is an implementation of the Livecoin API in Golang.
package livecoin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	API_BASE = "https://api.livecoin.net" // Livecoin API endpoint
)

// New returns an instantiated livecoin struct
func New(apiKey, apiSecret string) *Livecoin {
	client := NewClient(apiKey, apiSecret)
	return &Livecoin{client}
}

// NewWithCustomHttpClient returns an instantiated livecoin struct with custom http client
func NewWithCustomHttpClient(apiKey, apiSecret string, httpClient *http.Client) *Livecoin {
	client := NewClientWithCustomHttpConfig(apiKey, apiSecret, httpClient)
	return &Livecoin{client}
}

// NewWithCustomTimeout returns an instantiated livecoin struct with custom timeout
func NewWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) *Livecoin {
	client := NewClientWithCustomTimeout(apiKey, apiSecret, timeout)
	return &Livecoin{client}
}

// handleErr gets JSON response from livecoin API en deal with error
func handleErr(r interface{}) error {
	switch v := r.(type) {
	case map[string]interface{}:
		errorMessage := r.(map[string]interface{})["errorMessage"]
		if errorMessage != nil && errorMessage.(string) != "" {
			return errors.New(errorMessage.(string))
		}
	case []interface{}:
		return nil
	default:
		return fmt.Errorf("I don't know about type %T!\n", v)
	}

	return nil
}

// livecoin represent a livecoin client
type Livecoin struct {
	client *client
}

// set enable/disable http request/response dump
func (c *Livecoin) SetDebug(enable bool) {
	c.client.debug = enable
}

// Account

// GetBalances is used to retrieve all balances from your account
func (b *Livecoin) GetBalances() (balances []Balance, err error) {
	r, err := b.client.do("GET", "payment/balances", nil, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &balances)
	return
}

// Getbalance is used to retrieve the balance from your account for a specific currency.
// currency: a string literal for the currency (ex: LTC)
func (b *Livecoin) GetBalance(currency string) (balance Balance, err error) {
	r, err := b.client.do("GET", "payment/balance", map[string]string{"currency": strings.ToUpper(currency)}, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &balance)
	return
}

// GetOrderBook is used to retrieve the orderbook for specified currency pair
func (b *Livecoin) GetOrderBook(currency string) (orderbook Orderbook, err error) {
	r, err := b.client.do("GET", "exchange/order_book", map[string]string{"currencyPair": strings.ToUpper(currency)}, false)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &orderbook)
	return
}

// GetAllOrderBook is used to retrieve the orderbook for every currency pair.
func (b *Livecoin) GetAllOrderBook() (allOrderbook map[string]Orderbook, err error) {
	r, err := b.client.do("GET", "exchange/all/order_book", nil, false)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &allOrderbook)
	return
}

// GetTrades used to retrieve your trade history.
// market string literal for the market (ie. BTC/LTC). If set to "all", will return for all market
func (b *Livecoin) GetTrades(currencyPair string) (trades []Trade, err error) {
	payload := make(map[string]string)
	if currencyPair != "all" {
		payload["currencyPair"] = currencyPair
	}
	r, err := b.client.do("GET", "exchange/trades", payload, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	if res, ok := response.(map[string]interface{}); ok {
		if exception, ok := res["exception"]; ok && exception == "Data not found" {
			return
		}
	}
	err = json.Unmarshal(r, &trades)
	return
}

// GetTransactions is used to retrieve your withdrawal and deposit history
// "Start" and "end" are given in UNIX timestamp format in miliseconds and used to specify the date range for the data returned.
func (b *Livecoin) GetTransactions(start uint64, end uint64) (transactions []Transaction, err error) {
	if end == 0 {
		end = uint64(time.Now().Unix()) * 1000
	}
	r, err := b.client.do("GET", "payment/history/transactions", map[string]string{"types": "DEPOSIT,WITHDRAWAL", "start": strconv.FormatUint(uint64(start), 10), "end": strconv.FormatUint(uint64(end), 10)}, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &transactions)
	return
}

// SellLimit is used to place a sell order (limit) for a specified currency pair.
func (b *Livecoin) SellLimit(market string, quantity, rate decimal.Decimal) (order NewOrder, err error) {
	payload := map[string]string{
		"currencyPair": market,
		"price":        fmt.Sprintf("%s", rate),
		"quantity":     fmt.Sprintf("%s", quantity),
	}
	r, err := b.client.do("POST", "exchange/selllimit", payload, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &order)
	return
}

// BuyLimit is used to place a buy order (limit) for particular currency pair.
func (b *Livecoin) BuyLimit(market string, quantity, rate decimal.Decimal) (order NewOrder, err error) {
	payload := map[string]string{
		"currencyPair": market,
		"price":        fmt.Sprintf("%s", rate),
		"quantity":     fmt.Sprintf("%s", quantity),
	}
	r, err := b.client.do("POST", "exchange/buylimit", payload, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &order)
	return
}

// SellMarket is used to place a sell order (market) for specified amount of selected currency pair.
func (b *Livecoin) SellMarket(market string, quantity decimal.Decimal) (order NewOrder, err error) {
	payload := map[string]string{
		"currencyPair": market,
		"quantity":     fmt.Sprintf("%s", quantity),
	}
	r, err := b.client.do("POST", "exchange/sellmarket", payload, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &order)
	return
}

// BuyMarket is used to place a buy order (market) of specified amount for particular currency pair.
func (b *Livecoin) BuyMarket(market string, quantity decimal.Decimal) (order NewOrder, err error) {
	payload := map[string]string{
		"currencyPair": market,
		"quantity":     fmt.Sprintf("%s", quantity),
	}
	r, err := b.client.do("POST", "exchange/buymarket", payload, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &order)
	return
}

// CancelOrder is used to cancel an order. NOTE: market orders CANNOT be cancelled
func (b *Livecoin) CancelOrder(market, orderId string) (cancelledOrder CancelledOrder, err error) {
	payload := map[string]string{
		"currencyPair": market,
		"orderId":      orderId,
	}
	r, err := b.client.do("POST", "exchange/cancellimit", payload, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &cancelledOrder)
	return
}

// GetOrder is used to retrieve order information by its ID.
func (b *Livecoin) GetOrder(orderId string) (orderInfo OrderInfo, err error) {
	payload := map[string]string{
		"orderId": orderId,
	}
	r, err := b.client.do("GET", "exchange/order", payload, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &orderInfo)
	return
}

// GetRestrictions is used to retrieve the restrictions of all currency pairs
func (b *Livecoin) GetRestrictions() (restrictions Restrictions, err error) {
	r, err := b.client.do("GET", "exchange/restrictions", nil, false)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &restrictions)
	return
}
