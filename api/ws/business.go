package ws

import (
	"github.com/bitmyth/okex"
	"github.com/bitmyth/okex/events/public"
	requests "github.com/bitmyth/okex/requests/ws/public"
)

// Business
type Business struct {
	*ClientWs
	cCh    chan *public.Candlesticks
}

// NewPublic returns a pointer to a fresh Public
func NewBusiness(c *ClientWs) *Business {
	return &Business{ClientWs: c}
}

// Candlesticks
// Retrieve the open interest. Data will by pushed every 3 seconds.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-candlesticks-channel
func (c *Business) Candlesticks(req requests.Candlesticks, ch ...chan *public.Candlesticks) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.cCh = ch[0]
	}
	return c.Subscribe("business", []okex.ChannelName{okex.ChannelName(req.Channel)}, m)
}

// UCandlesticks
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-candlesticks-channel
func (c *Business) UCandlesticks(req requests.Candlesticks, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.cCh = nil
	}
	return c.Unsubscribe("business", []okex.ChannelName{okex.ChannelName(req.Channel)}, m)
}
