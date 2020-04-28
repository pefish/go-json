package go_json

import (
	"fmt"
	"testing"
)

func TestJsonClass_ParseToStruct(t *testing.T) {
	str := ` {"address":"AajrAMcZezLD6hEhuCxRkGtpuKUc6n2zKQ","amount":"1","coin":"NEO","confirmed_num":"1","confirming_threshold":"1","decimal":"0","id":"20190809171201322623000000012401","side":"deposit","status":"success","txid":"3d9c16fc2b62a66b077bde05f6fc777b42596e1cd638b3bfb2f712140920798e","vout_n":"0"}`
	type TxCallbackParam struct {
		Id                     string  `json:"id"`
		Coin                   string  `json:"coin"`
		Address                string  `json:"address"`
		Side                   string  `json:"side"`
		Amount                 string  `json:"amount"`
		Decimal                uint64  `json:"decimal"`
		TxId                   string  `json:"txid"`
		Vout                   uint64  `json:"vout_n"`
		Status                 string  `json:"status"`
		RequestId              string  `json:"request_id"`
		Confirmations          uint64  `json:"confirmed_num"`
		ConfirmationsThreshold uint64  `json:"confirming_threshold"`
		FeeCurrency            string  `json:"fee_coin"`
		FeeAmount              float64 `json:"fee_amount"`
		FeeDecimal             uint64  `json:"fee_decimal"`
		Type                   string  `json:"type"`
	}
	a := TxCallbackParam{}
	Json.ParseToStruct(str, &a)
	fmt.Printf(`%#v`, a)
}

func TestJsonClass_Stringify(t *testing.T) {
	a, err := Json.Stringify(map[interface{}]interface{}{"1": "1", 1: 2, false: 3,})
	if err != nil {
		t.Error()
	}
	if a != `{"1":"1","1":2,"false":3}` {
		t.Error()
	}
}
