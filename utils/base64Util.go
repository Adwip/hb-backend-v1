package utils

import "encoding/base64"

type base64Lib struct {
}

var RawStdEncoding = base64.StdEncoding.WithPadding(-1)

func Base64Lib() *base64Lib {
	base64 := &base64Lib{}
	return base64
}

func (base64Lib) Encode(data []byte) string {
	formatString := RawStdEncoding.EncodeToString(data)
	return formatString
}

func (base64Lib) Decode(data string) ([]byte, error) {
	result, err := RawStdEncoding.DecodeString(data)
	if err != nil {
		return result, err
	}
	return result, nil
}
