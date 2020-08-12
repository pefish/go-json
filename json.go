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

func (j *JsonClass) MustStringify(val interface{}) string {
	str, err := j.Stringify(val)
	if err != nil {
		panic(err)
	}
	return str
}

func (j *JsonClass) Stringify(val interface{}) (string, error) {
	result, err := j.Marshal(val)
	if err != nil {
		return ``, err
	}
	return string(result), nil
}

func (j *JsonClass) MustMarshal(val interface{}) []byte {
	result, err := j.Marshal(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (j *JsonClass) Marshal(val interface{}) ([]byte, error) {
	result, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (j *JsonClass) MustParse(str string) interface{} {
	result, err := j.Parse(str)
	if err != nil {
		panic(err)
	}
	return result
}

func (j *JsonClass) Parse(str string) (interface{}, error) {
	return j.ParseBytes([]byte(str))
}

func (j *JsonClass) MustParseToMap(str string) map[string]interface{} {
	map_, err := j.ParseToMap(str)
	if err != nil {
		panic(err)
	}
	return map_
}

func (j *JsonClass) ParseToMap(str string) (map[string]interface{}, error) {
	result, err := j.Parse(str)
	if err != nil {
		return nil, err
	}
	map_, ok := result.(map[string]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf(`%T cannot cast to map[string]interface{}`, result))
	}
	return map_, nil
}

func (j *JsonClass) MustParseBytesToMap(data []byte) map[string]interface{} {
	map_, err := j.ParseBytesToMap(data)
	if err != nil {
		panic(err)
	}
	return map_
}

func (j *JsonClass) ParseBytesToMap(data []byte) (map[string]interface{}, error) {
	result, err := j.ParseBytes(data)
	if err != nil {
		return nil, err
	}
	map_, ok := result.(map[string]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf(`%T cannot cast to map[string]interface{}`, result))
	}
	return map_, nil
}

func (j *JsonClass) MustParseBytes(bytes []byte) interface{} {
	result, err := j.ParseBytes(bytes)
	if err != nil {
		panic(err)
	}
	return result
}

func (j *JsonClass) ParseBytes(bytes []byte) (interface{}, error) {
	var result interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (j *JsonClass) MustParseToStruct(str string, struct_ interface{}) {
	err := j.ParseToStruct(str, struct_)
	if err != nil {
		panic(err)
	}
}

func (j *JsonClass) ParseToStruct(str string, struct_ interface{}) error {
	map_, err := j.ParseToMap(str)
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
