package utils

import "crypto/hmac"
import "encoding/hex"
import "crypto/sha256"
import _ "fmt"

type HashLib struct {
}

func Hash() *HashLib {
	hash := &HashLib{}
	return hash
}

func (HashLib) SHA256(data string, key string) string {
	hmacDeclare := hmac.New(sha256.New, []byte(key))
	hmacDeclare.Write([]byte(data))
	signature := hex.EncodeToString(hmacDeclare.Sum(nil))
	return signature
}

func (h HashLib) IsSHA256Valid(reqValue string, comparer string, key string) bool {
	result := h.SHA256(reqValue, key)
	return result == comparer
}

func (h HashLib) VerifyPassword(userInput, dbResult, passwordKey string) bool {
	hashed := h.SHA256(userInput, passwordKey)
	return hashed == dbResult
}
