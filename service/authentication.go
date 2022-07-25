package service

// import lib "hb-backend-v1/library"
import "hb-backend-v1/model/account"

// import "encoding/json"
import "fmt"

// import "os"
import "hb-backend-v1/repository"
import "github.com/gin-gonic/gin"

type AuthenticationInt interface {
	Login(*gin.Context) (bool, account.AuthResponse, string)
	LogOut()
}

type AuthenticationService struct {
	account repository.AccountInt
}

func NewAuthentication(account *repository.AccountInt) AuthenticationInt {
	return &AuthenticationService{
		account: *account,
	}
}

func (service AuthenticationService) Login(c *gin.Context) (bool, account.AuthResponse, string) {
	var authResponse account.AuthResponse
	// account := accountHandler.authentication
	var LoginForm account.LoginForm

	// var loginData accountForm.LoginData

	if err := c.ShouldBindJSON(&LoginForm); err != nil {
		// c.JSON(500, model.WebResponse{Success: false})
		// return
		fmt.Println(err)
	}
	fmt.Println(LoginForm)
	// fmt.Println(service.account.Login(c))
	/*
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
		}*/

	return true, authResponse, ""

}

func (AuthenticationService) LogOut() {

}
