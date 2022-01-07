package authentication

import "encoding/base64"
import "crypto/sha256"
import "encoding/json"
import "strings"
import "crypto/hmac"
import "encoding/hex"
import "fmt"




var key string = "12345678"
var payloadJson Payload
var RawStdEncoding = base64.StdEncoding.WithPadding(-1)

func GenerateToken(alg string, typ string, payload []byte)(string, error){
	var headerEncoded, payloadEncoded, signature, mergedEncoded string
	header := Header{Alg: alg,Typ: typ}


	headerJson, errHeader := json.Marshal(header)//struct -> byte
	if errHeader != nil{
		return "",errHeader
	}
	
	headerEncoded = encodeBase64(headerJson)//byte -> string
	payloadEncoded = encodeBase64(payload)//byte -> string
	mergedEncoded = headerEncoded+"."+payloadEncoded
	
	if alg=="SHA256"{
		signature = sha256encode(mergedEncoded, key)
	}

	finalToken := headerEncoded+"."+payloadEncoded+"."+signature
	
	// result, _ := VerifyToken(finalToken)
	// fmt.Printf("Token => %t", result)
	return finalToken, nil
}

func encodeBase64(data []byte) string {
	formatString := RawStdEncoding.EncodeToString(data)
	return formatString
}

func decodeBase64(data string)([]byte, error){
	result, err := RawStdEncoding.DecodeString(data)
	if err != nil{
		return result, err
	}
	return result, nil
}

func getHeader(header string){

}

func getPayload(){

}


func VerifyToken(token string)(bool, error){
	var headerJson Header
	split := strings.Split(token,".")
	if length := len(split); length != 3{
		return false, nil
	}

	headerDecoded,_	:= decodeBase64(split[0])
	headerErr		:= json.Unmarshal(headerDecoded, &headerJson)
	if headerErr != nil{
		fmt.Println(headerErr.Error())
	}
	// _				= headerErr
	payloadDecoded,_:= decodeBase64(split[1])
	payloadErr		:= json.Unmarshal(payloadDecoded, &payloadJson)
	if payloadErr != nil{
		fmt.Println(payloadErr.Error())
	}
	// _				= payloadErr
	signature			 	:= split[2]
	headerPayload	:= split[0]+"."+split[1]

	if alg := headerJson.Alg; alg == "SHA256"{
		result, _ := isSHA256KeyValid(headerPayload, signature)
		return result, nil
	}
	return false, nil
}

func isSHA256KeyValid(headerPayload string, signature string) (bool, error){
	var expectedSign = sha256encode(headerPayload, key)
	if expectedSign == signature {
		return true, nil
	}
	return false, nil
}

func sha256encode(data string, key string) string{
	var hmacDeclare = hmac.New(sha256.New, []byte(key))
	hmacDeclare.Write([]byte(data))
	var signature = hex.EncodeToString(hmacDeclare.Sum(nil))
	return signature
}

func compareSignature(message, signature, key []byte) bool{
	var hmacDeclare = hmac.New(sha256.New, key)
	hmacDeclare.Write(message)
	expectedMAC := hmacDeclare.Sum(nil)
	return hmac.Equal(signature, expectedMAC)
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