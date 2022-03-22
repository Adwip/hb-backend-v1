package library

import accountForm "hb-backend-v1/model/account"
import "encoding/base64"
import "encoding/json"
import _ "strings"
import "os"

type JWT struct {
}

var key string = os.Getenv("JWT_KEY")
var JWTPayload = accountForm.JWTPayload{}
var JWTHeader = accountForm.JWTHeader{}

func (jwt *JWT) GenerateToken(alg string, typ string, payload []byte) (string, error) {
	var headerEncoded, payloadEncoded, signature, mergedStringEncoded string
	var RawStdEncoding = base64.StdEncoding.WithPadding(-1)
	var cryptoEncode = Crypto{}

	header := accountForm.JWTHeader{Alg: alg, Typ: typ}

	headerJson, errHeader := json.Marshal(header)

	if errHeader != nil {
		return "", errHeader
	}

	headerEncoded = RawStdEncoding.EncodeToString(headerJson)
	payloadEncoded = RawStdEncoding.EncodeToString(payload)
	mergedStringEncoded = headerEncoded + "." + payloadEncoded

	if alg == "SHA256" {
		signature = cryptoEncode.SHA256(mergedStringEncoded, key)
	} else {
		signature = ""
	}

	finalToken := headerEncoded + "." + payloadEncoded + "." + signature

	return finalToken, nil
}

func (jwt *JWT) VerifiyToken(header string, payload string, signature string, headerObj accountForm.JWTHeader) bool {
	// mergedHeaderPayload := ""
	var encodedHeaderPayload string
	var cryptoEncode = Crypto{}
	mergedHeaderPayload := header + "." + payload
	if headerObj.Alg == "SHA256" {
		encodedHeaderPayload = cryptoEncode.SHA256(mergedHeaderPayload, key)
	} else {
		encodedHeaderPayload = ""
	}
	return encodedHeaderPayload == signature
}

func (jwt *JWT) DecodeToken(token string) (bool, accountForm.JWTHeader, accountForm.JWTPayload, error) {
	var header accountForm.JWTHeader
	var payload accountForm.JWTPayload

	return true, header, payload, nil
}
