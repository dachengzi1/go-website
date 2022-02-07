package utils

import "github.com/json-iterator/go"

func JsonMarshalToString(v interface{}) (string, error) {
	str, err := jsoniter.MarshalToString(v)
	return str, err
}
