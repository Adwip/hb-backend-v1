package utils

import m "hb-backend-v1/model"
import "encoding/json"
import "errors"
import "strings"
import _ "os"
import _ "fmt"

type JWT struct {
}

// var key string = os.Getenv("JWT_SECRET_KEY")
var JWTPayload m.JWTPayloadResponse
var JWTHeader m.JWTHeaderResponse

func JsonWT() *JWT {
	jwt := &JWT{}
	return jwt
}

func (jwt *JWT) GenerateToken(alg string, typ string, payload []byte, key string) (string, error) {
	var headerEncoded, payloadEncoded, signature, mergedStringEncoded string
	// RawStdEncoding := base64.StdEncoding.WithPadding(-1)
	cryptoEncode := Hash()
	base64 := Base64Lib()

	header := m.JWTHeaderResponse{Alg: alg, Typ: typ}

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

func (jwt *JWT) VerifiyToken(header string, payload string, signature string, headerObj m.JWTHeaderResponse, key string) bool {
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

func (jwt *JWT) DecodeToken(token string) (m.JWTHeaderResponse, m.JWTPayloadResponse, error) {
	var header m.JWTHeaderResponse
	var payload m.JWTPayloadResponse
	base64 := Base64Lib()
	split := strings.Split(token, ".")
	if length := len(split); length != 3 {
		return header, payload, errors.New("token not valid")
	}
	// fmt.Println("Token valid pass")

	headerBin, errHeader := base64.Decode(split[0])
	if errHeader != nil {
		// fmt.Println(split[0])
		return header, payload, errHeader
	}
	// fmt.Println("Decode split [0] pass")

	errHeader = json.Unmarshal(headerBin, &header)
	if errHeader != nil {
		return header, payload, errHeader
	}
	// fmt.Println("Unmarshall header pass")

	payloadBin, errPayload := base64.Decode(split[1])
	if errPayload != nil {
		return header, payload, errPayload
	}
	// fmt.Println("Decode split [1] pass")

	errPayload = json.Unmarshal(payloadBin, &payload)
	if errPayload != nil {
		return header, payload, errPayload
	}
	// fmt.Println("Unmarshall payload pass")

	return header, payload, nil
}
