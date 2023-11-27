package util

import "encoding/json"

func JsonEncode(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}
