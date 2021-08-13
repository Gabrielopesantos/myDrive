package utils

import (
	"bytes"
	"encoding/json"
)

// AnyToBytesBuffer Convert any object to Buffer helper
func AnyToBytesBuffer(i interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(i)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
