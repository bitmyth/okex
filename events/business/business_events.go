package business

import (
	"github.com/bitmyth/okex/events"
	"github.com/bitmyth/okex/models/market"
)

type (
	Candlesticks struct {
		Arg     *events.Argument        `json:"arg"`
		Candles []*market.CandleConfirm `json:"data"`
	}
)
