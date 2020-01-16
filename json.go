package go_json

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type JsonClass struct {

}

var Json = JsonClass{}

func (this *JsonClass) MustStringify(val interface{}) string {
	str, err := this.Stringify(val)
	if err != nil {
		panic(err)
	}
	return str
}

func (this *JsonClass) Stringify(val interface{}) (string, error) {
	result, err := json.Marshal(val)
	if err != nil {
		return ``, err
	}
	return string(result), nil
}

func (this *JsonClass) MustParse(str string) interface{} {
	result, err := this.Parse(str)
	if err != nil {
		panic(err)
	}
	return result
}

func (this *JsonClass) Parse(str string) (interface{}, error) {
	var result interface{}
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *JsonClass) MustParseToMap(str string) map[string]interface{} {
	map_, err := this.ParseToMap(str)
	if err != nil {
		panic(err)
	}
	return map_
}

func (this *JsonClass) ParseToMap(str string) (map[string]interface{}, error) {
	result, err := this.Parse(str)
	if err != nil {
		panic(err)
	}
	map_, ok := result.(map[string]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf(`%T cannot cast to map[string]interface{}`, result))
	}
	return map_, nil
}

func (this *JsonClass) MustParseBytes(bytes []byte) interface{} {
	result, err := this.ParseBytes(bytes)
	if err != nil {
		panic(err)
	}
	return result
}

func (this *JsonClass) ParseBytes(bytes []byte) (interface{}, error) {
	var result interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *JsonClass) MustParseToStruct(str string, struct_ interface{}) {
	err := this.ParseToStruct(str, struct_)
	if err != nil {
		panic(err)
	}
}

func (this *JsonClass) ParseToStruct(str string, struct_ interface{}) error {
	map_, err := this.ParseToMap(str)
	if err != nil {
		return err
	}
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		TagName:          "json",
		Result:           &struct_,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	err = decoder.Decode(map_)
	if err != nil {
		return err
	}
	return nil
}
