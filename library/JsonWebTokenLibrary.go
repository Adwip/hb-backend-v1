package library

import accountForm "hb-backend-v1/model/account"
import "encoding/json"
import "errors"
import "strings"
import _ "os"
import _ "fmt"

type JWT struct {
}

// var key string = os.Getenv("JWT_SECRET_KEY")
var JWTPayload accountForm.JWTPayload
var JWTHeader accountForm.JWTHeader

func JsonWT() *JWT {
	jwt := &JWT{}
	return jwt
}

func (jwt *JWT) GenerateToken(alg string, typ string, payload []byte, key string) (string, error) {
	var headerEncoded, payloadEncoded, signature, mergedStringEncoded string
	// RawStdEncoding := base64.StdEncoding.WithPadding(-1)
	cryptoEncode := Hash()
	base64 := Base64Lib()

	header := accountForm.JWTHeader{Alg: alg, Typ: typ}

	headerJson, errHeader := json.Marshal(header)

	if errHeader != nil {
		return "", errHeader
	}

	// headerEncoded = RawStdEncoding.EncodeToString(headerJson)
	// payloadEncoded = RawStdEncoding.EncodeToString(payload)
	headerEncoded = base64.Encode(headerJson)
	payloadEncoded = base64.Encode(payload)

	mergedStringEncoded = headerEncoded + "." + payloadEncoded

	if alg == "SHA256" {
		signature = cryptoEncode.SHA256(mergedStringEncoded, key)
	} else {
		signature = ""
	}

	finalToken := headerEncoded + "." + payloadEncoded + "." + signature

	return finalToken, nil
}

func (jwt *JWT) VerifiyToken(header string, payload string, signature string, headerObj accountForm.JWTHeader, key string) bool {
	// mergedHeaderPayload := ""
	cryptoEncode := Hash()
	var encodedHeaderPayload string
	mergedHeaderPayload := header + "." + payload
	if headerObj.Alg == "SHA256" {
		encodedHeaderPayload = cryptoEncode.SHA256(mergedHeaderPayload, key)
	} else {
		encodedHeaderPayload = ""
	}
	return encodedHeaderPayload == signature
}

func (jwt *JWT) DecodeToken(token string) (accountForm.JWTHeader, accountForm.JWTPayload, error) {
	var header accountForm.JWTHeader
	var payload accountForm.JWTPayload
	base64 := Base64Lib()
	split := strings.Split(token, ".")
	if length := len(split); length != 3 {
		return header, payload, errors.New("token not valid")
	}

	headerBin, errHeader := base64.Decode(split[0])
	if errHeader != nil {
		return header, payload, errHeader
	}
	errHeader = json.Unmarshal(headerBin, &header)
	if errHeader != nil {
		return header, payload, errHeader
	}

	payloadBin, errPayload := base64.Decode(split[1])
	if errPayload != nil {
		return header, payload, errPayload
	}
	errPayload = json.Unmarshal(payloadBin, &payload)
	if errPayload != nil {
		return header, payload, errPayload
	}

	return header, payload, nil
}
