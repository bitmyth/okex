package ws

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bitmyth/okex"
	"github.com/bitmyth/okex/events"
	"github.com/bitmyth/okex/events/business"
	requests "github.com/bitmyth/okex/requests/ws/public"
)

// Business
type Business struct {
	*ClientWs
	cCh chan *business.Candlesticks
}

// NewPublic returns a pointer to a fresh Public
func NewBusiness(c *ClientWs) *Business {
	return &Business{ClientWs: c}
}

// Candlesticks
// Retrieve the open interest. Data will by pushed every 3 seconds.
//
// https://www.okex.com/docs-v5/en/#websocket-api-public-channels-candlesticks-channel
func (c *Business) Candlesticks(req requests.Candlesticks, ch ...chan *business.Candlesticks) error {
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

// {"arg":{"channel":"candle1m","instId":"BTC-USDT-SWAP"},"data":[["1774964340000","66898.7","66902.1","66884.6","66894.5","1330.07","13.3007","889750.68514000","0"]]}

func (c *Business) Process(data []byte, e *events.Basic) bool {
	if e.Event == "" {
		ch, ok := e.Arg.Get("channel")
		if !ok {
			return false
		}
		switch ch {
		default:
			// special cases
			// market price candlestick channel
			chName := fmt.Sprint(ch)
			// candlestick channels
			if strings.Contains(chName, "candle") {
				e := business.Candlesticks{}
				err := json.Unmarshal(data, &e)
				if err != nil {
					println(err.Error())
					return false
				}
				go func() {
					if c.cCh != nil {
						c.cCh <- &e
					}
					c.StructuredEventChan <- e
				}()
				return true
			}
		}
	}
	return false
}
