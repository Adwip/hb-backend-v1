package service

import lib "hb-backend-v1/library"
import "hb-backend-v1/model/account"
import "encoding/json"
import "fmt"
import "os"

type authentication struct {
}

func Auth() *authentication {
	return &authentication{}
}

func (authentication) CreateLoginSession(req account.LoginData) (bool, account.AuthResponse, string) {
	var authResponse account.AuthResponse
	jwtLib := lib.JsonWT()
	currentDateTime := lib.Time().CurrentTimeUnix()
	jwtKey := os.Getenv("JWT_SECRET_KEY")

	JWTPayload := account.JWTPayload{
		AccountID:      req.AccountID,
		UserID:         req.UserID,
		CustomerID:     req.CustomerID,
		FirstName:      req.FirstName,
		PrimaryAccount: req.PrimaryAccount,
		AccountStatus:  req.AccountStatus,
		TimeZone:       req.TimeZone,
		CreatedAt:      currentDateTime,
	}

	payload, errJson := json.Marshal(JWTPayload)
	if errJson != nil {
		fmt.Println(errJson)
		return false, authResponse, "Failed to generate token"
	}

	token, errToken := jwtLib.GenerateToken("SHA256", "JWT", payload, jwtKey)
	if errToken != nil {
		fmt.Println(errToken)
		return false, authResponse, "Login rejected"
	}

	authResponse = account.AuthResponse{
		AccountID:      req.AccountID,
		UserID:         req.UserID,
		CustomerID:     req.CustomerID,
		FirstName:      req.FirstName,
		PrimaryAccount: req.PrimaryAccount,
		AccountStatus:  req.AccountStatus,
		TimeZone:       req.TimeZone,
		CreatedAt:      currentDateTime,
		Token:          token,
	}

	return true, authResponse, ""

}

func (authentication) DestryoLoginSession() {

}
