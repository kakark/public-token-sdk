package helper

import (
	"encoding/base64"
	"encoding/json"
)

func DecodeAndUnmarshal(part string, target interface{}) error {
	decodeBytes, err := base64.RawURLEncoding.DecodeString(part)
	if err != nil {
		return err
	}
	err = json.Unmarshal(decodeBytes, target)
	if err != nil {
		return err
	}
	return nil
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(s string) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return decodeBytes, nil
}
