package go_json

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
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

func (this *JsonClass) ParseBytes(bytes []byte) interface{} {
	var result interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		panic(err)
	}
	return result
}

func (this *JsonClass) ParseToStruct(str string, struct_ interface{}) {
	map_ := this.ParseToMap(str)
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		TagName:          "json",
		Result:           &struct_,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(map_)
	if err != nil {
		panic(err)
	}
}
