package controller

import "hb-backend-v1/model"
import "github.com/gin-gonic/gin"
import "hb-backend-v1/service"

type AccountCtrl struct {
	authenticationService service.AuthenticationInt
}

func Account(auth *service.AuthenticationInt) *AccountCtrl {
	return &AccountCtrl{
		authenticationService: *auth,
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
func (handler AccountCtrl) Regristration(c *gin.Context) {
	var RegistrationForm model.RegistrationRequest

	if err := c.ShouldBindJSON(&RegistrationForm); err != nil {
		c.JSON(400, gin.H{"success": false, "message": err.Error()})
		return
	}

	success, id, msg := handler.authenticationService.Registration(c, RegistrationForm)

	if !success {
		c.JSON(200, model.WebResponse{Success: false, Msg: msg})
		return
	}
	c.JSON(200, model.WebResponse{Success: true, Data: id})
}

/*

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
