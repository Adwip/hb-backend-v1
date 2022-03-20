package library

import "crypto/hmac"
import "encoding/hex"
import "crypto/sha256"

type Crypto struct {
}

func (c *Crypto) SHA256(data string, key string) string {
	hmacDeclare := hmac.New(sha256.New, []byte(key))
	hmacDeclare.Write([]byte(data))
	signature := hex.EncodeToString(hmacDeclare.Sum(nil))
	return signature
}
