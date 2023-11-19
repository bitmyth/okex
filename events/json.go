package events

import (
	"bytes"
	"encoding/json"
	"github.com/bitmyth/okex"
	"io"
)

func DecodeBasicEvent(data []byte) (e Basic, err error) {
	dec := json.NewDecoder(bytes.NewReader(data))
	var t json.Token

	for {
		t, err = dec.Token()
		if err != nil && err != io.EOF {
			return
		}
		if err == io.EOF {
			err = nil
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

		if t == "arg" {
			for {
				t, err = dec.Token()
				if err != nil && err != io.EOF {
					return
				}
				if t == nil {
					break
				}

				if t == "arg" {
					for dec.More() {
						t, err = dec.Token()
						if err != nil {
							return
						}

						if t == "channel" {

							var p string
							err = dec.Decode(&p)
							if err != nil {
								return
							}

							e.Arg = &Argument{
								arg: map[string]interface{}{
									"channel": p,
								},
							}
						}

					}

				}
			}

			var p string
			err = dec.Decode(&p)
			e.Msg = p
		}
	}

	return
}
