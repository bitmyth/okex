package events

import (
	"encoding/json"
	"testing"
)

func TestDecodeBasicEvent(t *testing.T) {
	data := []byte(`{"event":"foo", "arg": {"arg":{"channel":"bar"} }}`)
	gotE, err := DecodeBasicEvent(data)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%#v", gotE)
	marshal, err := json.Marshal(gotE)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(marshal))
}

type Argument2 struct {
	Arg map[string]interface{} `json:"arg"`
}

func Test_jxx(t *testing.T) {
	var arg Argument2
	arg.Arg = make(map[string]interface{})
	arg.Arg["a"] = "33"

	marshal, err := json.Marshal(arg)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(marshal))

}
