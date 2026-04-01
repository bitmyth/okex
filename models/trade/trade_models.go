package trade

import (
	"github.com/bitmyth/okex"
)

type (
	PlaceOrder struct {
		ClOrdID string         `json:"clOrdId"`
		Tag     string         `json:"tag"`
		SMsg    string         `json:"sMsg"`
		SCode   okex.JSONInt64 `json:"sCode"`
		OrdID   string         `json:"ordId"`
	}
	CancelOrder struct {
		OrdID   string           `json:"ordId"`
		ClOrdID string           `json:"clOrdId"`
		SMsg    string           `json:"sMsg"`
		SCode   okex.JSONFloat64 `json:"sCode"`
	}
	AmendOrder struct {
		OrdID   string           `json:"ordId"`
		ClOrdID string           `json:"clOrdId"`
		ReqID   string           `json:"reqId"`
		SMsg    string           `json:"sMsg"`
		SCode   okex.JSONFloat64 `json:"sCode"`
	}
	ClosePosition struct {
		InstID  string            `json:"instId"`
		PosSide okex.PositionSide `json:"posSide"`
	}
	Order struct {
		InstID         string              `json:"instId"`
		Ccy            string              `json:"ccy"`
		OrdID          string              `json:"ordId"`
		ClOrdID        string              `json:"clOrdId"`
		TradeID        string              `json:"tradeId"`
		Tag            string              `json:"tag"`
		Category       string              `json:"category"`
		FeeCcy         string              `json:"feeCcy"`
		RebateCcy      string              `json:"rebateCcy"`
		Px             okex.JSONFloat64    `json:"px"`
		Sz             okex.JSONFloat64    `json:"sz"`
		Pnl            okex.JSONFloat64    `json:"pnl"`
		AccFillSz      okex.JSONFloat64    `json:"accFillSz"`
		FillPx         okex.JSONFloat64    `json:"fillPx"`
		FillSz         okex.JSONFloat64    `json:"fillSz"`
		FillTime       okex.JSONFloat64    `json:"fillTime"`
		AvgPx          okex.JSONFloat64    `json:"avgPx"`
		Lever          okex.JSONFloat64    `json:"lever"`
		TpTriggerPx    okex.JSONFloat64    `json:"tpTriggerPx"`
		TpOrdPx        okex.JSONFloat64    `json:"tpOrdPx"`
		SlTriggerPx    okex.JSONFloat64    `json:"slTriggerPx"`
		SlOrdPx        okex.JSONFloat64    `json:"slOrdPx"`
		Fee            okex.JSONFloat64    `json:"fee"`
		Rebate         okex.JSONFloat64    `json:"rebate"`
		State          okex.OrderState     `json:"state"`
		TdMode         okex.TradeMode      `json:"tdMode"`
		PosSide        okex.PositionSide   `json:"posSide"`
		Side           okex.OrderSide      `json:"side"`
		OrdType        okex.OrderType      `json:"ordType"`
		InstType       okex.InstrumentType `json:"instType"`
		TgtCcy         okex.QuantityType   `json:"tgtCcy"`
		UTime          okex.JSONTime       `json:"uTime"`
		CTime          okex.JSONTime       `json:"cTime"`
		AttachAlgoOrds []AttachAlgoOrder   `json:"attachAlgoOrds"`
	}

	TransactionDetail struct {
		InstID   string              `json:"instId"`
		OrdID    string              `json:"ordId"`
		TradeID  string              `json:"tradeId"`
		ClOrdID  string              `json:"clOrdId"`
		BillID   string              `json:"billId"`
		Tag      okex.JSONFloat64    `json:"tag"`
		FillPx   okex.JSONFloat64    `json:"fillPx"`
		FillSz   okex.JSONFloat64    `json:"fillSz"`
		FeeCcy   okex.JSONFloat64    `json:"feeCcy"`
		Fee      okex.JSONFloat64    `json:"fee"`
		InstType okex.InstrumentType `json:"instType"`
		Side     okex.OrderSide      `json:"side"`
		PosSide  okex.PositionSide   `json:"posSide"`
		ExecType okex.OrderFlowType  `json:"execType"`
		TS       okex.JSONTime       `json:"ts"`
	}
	PlaceAlgoOrder struct {
		AlgoID string         `json:"algoId"`
		SMsg   string         `json:"sMsg"`
		SCode  okex.JSONInt64 `json:"sCode"`
	}
	CancelAlgoOrder struct {
		AlgoID string         `json:"algoId"`
		SMsg   string         `json:"sMsg"`
		SCode  okex.JSONInt64 `json:"sCode"`
	}
	AlgoOrder struct {
		InstID       string              `json:"instId"`
		Ccy          string              `json:"ccy"`
		OrdID        string              `json:"ordId"`
		AlgoID       string              `json:"algoId"`
		ClOrdID      string              `json:"clOrdId"`
		TradeID      string              `json:"tradeId"`
		Tag          string              `json:"tag"`
		Category     string              `json:"category"`
		FeeCcy       string              `json:"feeCcy"`
		RebateCcy    string              `json:"rebateCcy"`
		TimeInterval string              `json:"timeInterval"`
		Px           okex.JSONFloat64    `json:"px"`
		PxVar        okex.JSONFloat64    `json:"pxVar"`
		PxSpread     okex.JSONFloat64    `json:"pxSpread"`
		PxLimit      okex.JSONFloat64    `json:"pxLimit"`
		Sz           okex.JSONFloat64    `json:"sz"`
		SzLimit      okex.JSONFloat64    `json:"szLimit"`
		ActualSz     okex.JSONFloat64    `json:"actualSz"`
		ActualPx     okex.JSONFloat64    `json:"actualPx"`
		Pnl          okex.JSONFloat64    `json:"pnl"`
		AccFillSz    okex.JSONFloat64    `json:"accFillSz"`
		FillPx       okex.JSONFloat64    `json:"fillPx"`
		FillSz       okex.JSONFloat64    `json:"fillSz"`
		FillTime     okex.JSONFloat64    `json:"fillTime"`
		AvgPx        okex.JSONFloat64    `json:"avgPx"`
		Lever        okex.JSONFloat64    `json:"lever"`
		TpTriggerPx  okex.JSONFloat64    `json:"tpTriggerPx"`
		TpOrdPx      okex.JSONFloat64    `json:"tpOrdPx"`
		SlTriggerPx  okex.JSONFloat64    `json:"slTriggerPx"`
		SlOrdPx      okex.JSONFloat64    `json:"slOrdPx"`
		OrdPx        okex.JSONFloat64    `json:"ordPx"`
		Fee          okex.JSONFloat64    `json:"fee"`
		Rebate       okex.JSONFloat64    `json:"rebate"`
		State        okex.OrderState     `json:"state"`
		TdMode       okex.TradeMode      `json:"tdMode"`
		ActualSide   okex.PositionSide   `json:"actualSide"`
		PosSide      okex.PositionSide   `json:"posSide"`
		Side         okex.OrderSide      `json:"side"`
		OrdType      okex.AlgoOrderType  `json:"ordType"`
		InstType     okex.InstrumentType `json:"instType"`
		TgtCcy       okex.QuantityType   `json:"tgtCcy"`
		CTime        okex.JSONTime       `json:"cTime"`
		TriggerTime  okex.JSONTime       `json:"triggerTime"`
	}

	AttachAlgoOrder struct {
		AttachAlgoId      string `json:"attachAlgoId,omitempty"`      // The order ID of attached TP/SL order (用于 amend)
		AttachAlgoClOrdId string `json:"attachAlgoClOrdId,omitempty"` // Client-supplied Algo ID

		// Take Profit (TP) 相关字段
		TpOrdKind       string `json:"tpOrdKind,omitempty"`       // condition / limit
		TpTriggerPx     string `json:"tpTriggerPx,omitempty"`     // Take-profit trigger price
		TpTriggerRatio  string `json:"tpTriggerRatio,omitempty"`  // Take profit trigger ratio (e.g. "0.3" = 30%)
		TpTriggerPxType string `json:"tpTriggerPxType,omitempty"` // last / index / mark
		TpOrdPx         string `json:"tpOrdPx,omitempty"`         // Take-profit order price

		// Stop Loss (SL) 相关字段
		SlTriggerPx     string `json:"slTriggerPx,omitempty"`     // Stop-loss trigger price
		SlTriggerRatio  string `json:"slTriggerRatio,omitempty"`  // Stop-loss trigger ratio
		SlTriggerPxType string `json:"slTriggerPxType,omitempty"` // last / index / mark
		SlOrdPx         string `json:"slOrdPx,omitempty"`         // Stop-loss order price

		// 其他字段
		Sz                   string `json:"sz,omitempty"`                   // Size (仅适用于 split TP)
		AmendPxOnTriggerType string `json:"amendPxOnTriggerType,omitempty"` // 是否启用成本价止损 (0/1)，仅适用于 split TP 的 SL
		FailCode             string `json:"failCode,omitempty"`             // 失败时的错误码
		FailReason           string `json:"failReason,omitempty"`           // 失败原因
	}
)
