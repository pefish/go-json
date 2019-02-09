package p_json

import (
	"encoding/json"
)

type JsonClass struct {

}

var Json = JsonClass{}

func (this *JsonClass) Stringify(val interface{}) string {
	result, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	return string(result)
}

func (this *JsonClass) Parse(str string) interface{} {
	var result interface{}
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		panic(err)
	}
	return result
}

func (this *JsonClass) ParseWithErr(str string) (interface{}, error) {
	var result interface{}
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *JsonClass) ParseToMap(str string) map[string]interface{} {
	var result interface{}
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		panic(err)
	}
	return result.(map[string]interface{})
}

func (this *JsonClass) ParseToMapWithErr(str string) (map[string]interface{}, error) {
	var result interface{}
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		return nil, err
	}
	return result.(map[string]interface{}), nil
}

func (this *JsonClass) ParseBytesToMapString(bytes []byte) map[string]string {
	var result interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		panic(err)
	}
	temp := result.(map[string]interface{})
	out := map[string]string{}
	for key, value := range temp {
		out[key] = value.(string)
	}
	return out
}

func (this *JsonClass) ParseBytes(bytes []byte) interface{} {
	var result interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		panic(err)
	}
	return result
}

func (this *JsonClass) ParseToStruct(str string, struct_ interface{}) {
	if err := json.Unmarshal([]byte(str), &struct_); err != nil {
		panic(err)
	}
}

func (this *JsonClass) ParseBytesToStruct(bytes []byte, struct_ interface{}) {
	if err := json.Unmarshal(bytes, &struct_); err != nil {
		panic(err)
	}
}
