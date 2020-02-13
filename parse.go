package js

import (
	"encoding/json"
	"fmt"
)

func Parse(data []byte) (Value, error) {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}

	var recursively func(v interface{}) Value

	recursively = func(v interface{}) Value {
		switch x := v.(type) {
		case map[string]interface{}:
			out := make(Object, len(x))
			for k, v := range x {
				out[k] = recursively(v)
			}
			return out
		case []interface{}:
			out := make(Array, len(x))
			for i := range x {
				out[i] = recursively(x[i])
			}
			return out
		case string:
			return String(x)
		case bool:
			return Bool(x)
		case nil:
			return Null{}
		case float64:
			return Number(x)
		default:
			panic(fmt.Sprintf("unknown value: %+v", x))
		}
	}

	return recursively(v), nil
}

func ParseObject(data []byte) (Object, error) {
	v, err := Parse(data)
	if err != nil {
		return nil, err
	}

	return v.Object(), nil
}

func ParseArray(data []byte) (Array, error) {
	v, err := Parse(data)
	if err != nil {
		return nil, err
	}

	return v.Array(), nil
}
