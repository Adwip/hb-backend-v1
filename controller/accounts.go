package controller

import "hb-backend-v1/model"

// import "hb-backend-v1/repository"
// import "github.com/gin-gonic/gin"
// import accountForm "hb-backend-v1/model/account"
import "github.com/gin-gonic/gin"

// import accountForm "hb-backend-v1/model/account"
import "hb-backend-v1/service"

// import "reflect"

type AccountCtrl struct {
	authenticationService service.AuthenticationInt
}

func Account(auth *service.AuthenticationInt) *AccountCtrl {
	return &AccountCtrl{
		authenticationService: *auth,
	}
}

func (accountHandler AccountCtrl) Routes(router *gin.Engine) {
	routes := router.Group("/auth")
	{
		routes.POST("/", accountHandler.Login)
	}
}

func (handler AccountCtrl) Login(c *gin.Context) {
	var loginForm model.LoginRequest

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.JSON(500, model.WebResponse{Success: false})
		return
	}

	success, result, msg := handler.authenticationService.Login(c, loginForm)

	if !success {
		c.JSON(200, model.WebResponse{Success: false, Msg: msg})
		return
	}
	c.JSON(200, model.WebResponse{Success: true, Data: result})
}

/*
func (AccountCtrl) Regristration(c *gin.Context) {
	account := repository.Account()
	var RegistrationForm accountForm.RegistrationForm

	if err := c.ShouldBindJSON(&RegistrationForm); err != nil {
		c.JSON(400, gin.H{"success": false, "message": err.Error()})
		return
	}
	result := account.RegistrationUser(c, RegistrationForm)

	if result.Success {
		c.JSON(200, model.WebResponse{Success: true})
		return
	}
	c.JSON(200, model.WebResponse{Success: false, Msg: result.Msg})
}

func (AccountCtrl) UpdatePassword(c *gin.Context) {
	var UpdatePasswordForm accountForm.UpdatePasswordForm
	account := repository.Account()
	if err := c.ShouldBindJSON(&UpdatePasswordForm); err != nil {
		c.JSON(400, gin.H{"success": false})
		return
	}
	result := account.UpdatePassword(c, UpdatePasswordForm)
	if result.Success {
		c.JSON(200, model.WebResponse{Success: true})
		return
	}
	c.JSON(200, model.WebResponse{Success: true, Msg: result.Msg})
}*/

/*
func Test(c *gin.Context) {
	c.JSON(200, gin.H{"success": true, "result": "This is test"})
}*/
