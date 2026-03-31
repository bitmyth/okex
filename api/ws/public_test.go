package ws

import (
	"github.com/bitmyth/okex/events/public"
	"testing"
)

var s = `
[
  {
    "DepthPrice": 6.308,
    "Size": 2.103,
    "LiquidatedOrder": 0,
    "OrderNumbers": 1
  },
  {
    "DepthPrice": 6.309,
    "Size": 0,
    "LiquidatedOrder": 0,
    "OrderNumbers": 0
  },
  {
    "DepthPrice": 6.31,
    "Size": 91.5016,
    "LiquidatedOrder": 0,
    "OrderNumbers": 3
  },
  {
    "DepthPrice": 6.323,
    "Size": 703.684,
    "LiquidatedOrder": 0,
    "OrderNumbers": 6
  },
  {
    "DepthPrice": 6.325,
    "Size": 340.2033,
    "LiquidatedOrder": 0,
    "OrderNumbers": 4
  },
  {
    "DepthPrice": 7.55,
    "Size": 1.0232,
    "LiquidatedOrder": 0,
    "OrderNumbers": 1
  }
]
`

var mp = `
{
  "arg": {},
  "data": [
    {
      "instId": "TIA-USDT",
      "instType": "MARGIN",
      "markPx": 6.674,
      "ts": {}
    }
  ]
}
`

var mpNull = `
{
  "arg": {},
  "data":null 
}
`

func Test_decodeMarketPrice(t *testing.T) {
	var price public.MarkPrice
	data := []byte(mpNull)
	if err := decodeMarketPrice(data, &price); err != nil {
		t.Error(err)
	}
}
