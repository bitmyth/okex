package decoder

import (
	"bytes"
	"encoding/json"
	"github.com/bitmyth/okex"
	"github.com/bitmyth/okex/events"
	"io"
)

func DecodeBasicEvent(data []byte) (e events.Basic, err error) {
	dec := json.NewDecoder(bytes.NewReader(data))
	var t json.Token

	for {
		t, err = dec.Token()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}
		if t == nil {
			break
		}

		if t == "event" {
			var p string
			err = dec.Decode(&p)
			if err != nil {
				return
			}
			e.Event = p
		}

		if t == "op" {
			var p string
			err = dec.Decode(&p)
			if err != nil {
				return
			}
			e.Op = okex.Operation(p)
		}

		if t == "code" {
			var p int
			err = dec.Decode(&p)
			if err != nil {
				return
			}
			e.Code = p
		}

		if t == "msg" {
			var p string
			err = dec.Decode(&p)
			if err != nil {
				return
			}
			e.Msg = p
		}
	}
	return
}
