package auth

import "crypto/md5"
import "encoding/hex"
import _"fmt"

func VerifyPassword(userInput string, userData string) bool{
	hasher := md5.New()
	hasher.Write([]byte(userInput))
	
	userInput = hex.EncodeToString(hasher.Sum(nil))

	if userInput == userData{
		return true
	}

	return false
}

func VerifyJWT(){

}