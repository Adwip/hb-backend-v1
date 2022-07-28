package service

import "hb-backend-v1/utils"
import "hb-backend-v1/model"

import "encoding/json"
import "fmt"

import "os"
import "hb-backend-v1/repository"
import "github.com/gin-gonic/gin"

type AuthenticationInt interface {
	Login(*gin.Context) (bool, model.AccountLoginResponse, string)
	LogOut()
}

type AuthenticationService struct {
	accountRepo repository.AccountInt
}

func NewAuthentication(account *repository.AccountInt) AuthenticationInt {
	return &AuthenticationService{
		accountRepo: *account,
	}
}

func (service AuthenticationService) Login(c *gin.Context) (bool, model.AccountLoginResponse, string) {
	var loginResult model.AccountLoginResponse
	// account := accountHandler.authentication
	var loginForm model.LoginRequest

	// var loginData accountForm.LoginData

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		// c.JSON(500, model.WebResponse{Success: false})
		// return
		fmt.Println(err)
	}

	exists, accountData, msg := service.accountRepo.Login(c, loginForm)

	if !exists {
		return false, loginResult, msg
	}

	// fmt.Println(service.account.Login(c))

	jwtLib := utils.JsonWT()
	currentDateTime := utils.Time().CurrentTimeUnix()
	jwtKey := os.Getenv("JWT_SECRET_KEY")

	JWTPayload := model.JWTPayloadResponse{
		AccountID:      accountData.AccountID,
		UserID:         accountData.UserID,
		CustomerID:     accountData.CustomerID,
		FirstName:      accountData.FirstName,
		PrimaryAccount: accountData.PrimaryAccount,
		AccountStatus:  accountData.AccountStatus,
		TimeZone:       accountData.TimeZone,
		CreatedAt:      currentDateTime,
	}

	payload, errJson := json.Marshal(JWTPayload)
	if errJson != nil {
		fmt.Println(errJson)
		return false, loginResult, "Failed to generate token"
	}

	token, errToken := jwtLib.GenerateToken("SHA256", "JWT", payload, jwtKey)
	if errToken != nil {
		fmt.Println(errToken)
		return false, loginResult, "Login rejected"
	}

	loginResult = model.AccountLoginResponse{
		AccountID:      accountData.AccountID,
		UserID:         accountData.UserID,
		CustomerID:     accountData.CustomerID,
		FirstName:      accountData.FirstName,
		PrimaryAccount: accountData.PrimaryAccount,
		AccountStatus:  accountData.AccountStatus,
		TimeZone:       accountData.TimeZone,
		CreatedAt:      currentDateTime,
		Token:          token,
	}

	return true, loginResult, ""

}

func (AuthenticationService) LogOut() {

}
