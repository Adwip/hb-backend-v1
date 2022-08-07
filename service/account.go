package service

import "hb-backend-v1/utils"
import "hb-backend-v1/model"
import "encoding/json"
import "fmt"
import "os"
import "hb-backend-v1/repository"
import "github.com/gin-gonic/gin"

type Account interface {
	Registration(*gin.Context, model.RegistrationRequest) (bool, string, string)
	Login(*gin.Context, model.LoginRequest) (bool, model.AccountLoginResponse, string)
	LogOut()
	UpdatePassword(*gin.Context, *model.UpdatePasswordRequest) (bool, string)
}

type AccountService struct {
	accountRepo repository.Account
}

func NewAccountService(account *repository.Account) Account {
	return &AccountService{
		accountRepo: *account,
	}
}

func (service AccountService) Registration(c *gin.Context, req model.RegistrationRequest) (bool, string, string) {
	hash := utils.Hash()
	passwordKey := os.Getenv("PASSWORD_SECRET_KEY")
	req.Password = hash.SHA256(req.Password, passwordKey)

	success, id, msg := service.accountRepo.Registration(c, req)
	return success, id, msg
}

func (service AccountService) Login(c *gin.Context, loginForm model.LoginRequest) (bool, model.AccountLoginResponse, string) {
	var loginResult model.AccountLoginResponse

	exists, accountData, msg := service.accountRepo.Login(c, loginForm)

	if !exists {
		return false, loginResult, msg
	}

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

func (service AccountService) UpdatePassword(c *gin.Context, req *model.UpdatePasswordRequest) (bool, string) {
	hash := utils.Hash()
	passwordKey := os.Getenv("PASSWORD_SECRET_KEY")
	req.NewPassword = hash.SHA256(req.NewPassword, passwordKey)
	req.OldPassword = hash.SHA256(req.OldPassword, passwordKey)
	req.ConfirmPassword = hash.SHA256(req.ConfirmPassword, passwordKey)

	if req.ConfirmPassword != req.NewPassword {
		return false, "Password Not Matched"
	}

	success, msg := service.accountRepo.UpdatePassword(c, req)

	if !success {
		return false, msg
	}
	return true, "Succes update password"
}

func (AccountService) LogOut() {

}
