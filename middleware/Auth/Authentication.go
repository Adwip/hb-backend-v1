package auth

import "encoding/base64"
import "crypto/sha256"
import "encoding/json"
import "strings"
import "crypto/hmac"
import "encoding/hex"
import "fmt"

type Header struct{
	Alg string
	Typ string
}

var key string = "12345678"

func GenerateToken(alg string, typ string, payload []byte)(string, error){
	var headerString, payloadString, signatureString, mergedString string
	header := Header{Alg: alg,Typ: typ}
	hmacDeclare := hmac.New(sha256.New, []byte(key))


	headerJson, errHeader := json.Marshal(header)
	if errHeader != nil{
		return "",errHeader
	}
	
	headerString = toBase64(headerJson)
	payloadString = toBase64(payload)
	mergedString = headerString+"."+payloadString
	if alg=="SHA256"{
		hmacDeclare.Write([]byte(mergedString))
		signatureString = hex.EncodeToString(hmacDeclare.Sum(nil))
	}

	finalToken := headerString+"."+payloadString+"."+signatureString
	return finalToken, nil
}

func toBase64(data []byte) string {
	formatString := strings.TrimRight(base64.StdEncoding.EncodeToString(data),"=")
	return formatString
}

func getJWT(){

}

func parsingJWT(){

}

func GetName(){

}

func GetID(){

}

func GetUserType(){

}


func VerifyToken(token string)(bool, error){
	return true, nil
}