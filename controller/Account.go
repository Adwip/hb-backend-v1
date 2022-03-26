package controller

import "hb-backend-v1/model"
import "hb-backend-v1/repository"
import "github.com/gin-gonic/gin"
import accountForm "hb-backend-v1/model/account"
import "hb-backend-v1/library"

func AllAccount(c *gin.Context) {
	identity := library.Identity(c)
	c.JSON(200, model.WebResponse{Success: true, Data: identity.GetUserID()})

	/*
		result, err := account.AllAccount()

		if err != nil {
			c.JSON(200, gin.H{"hasil": err.Error()})
		} else {
			c.JSON(200, gin.H{"success": true, "result": result})
		}*/
}

func Login(c *gin.Context) {
	account := repository.Account()
	var LoginForm accountForm.LoginForm

	if err := c.ShouldBindJSON(&LoginForm); err != nil {
		c.JSON(500, model.WebResponse{Success: false})
		return
	}

	result := account.Login(c, &LoginForm)
	if result.Success {
		c.JSON(200, model.WebResponse{Success: true, Data: result.Data})
		return
	}
	c.JSON(200, model.WebResponse{Success: false, Msg: result.Msg})

}

/*
func Regristration(c *gin.Context) {
	var RegistrationForm accountForm.RegistrationForm

	if err := c.ShouldBindJSON(&RegistrationForm); err != nil {
		c.JSON(400, gin.H{"success": false, "message": err.Error()})
		return
	}
	success, err := account.RegistrationUser(RegistrationForm)
	if success {
		c.JSON(200, gin.H{"success": true, "data": RegistrationForm})
		return
	}
	c.JSON(400, gin.H{"success": false, "message": err.Error()})
}*/

/*
func UpdatePassword(c *gin.Context){
	var UpdatePasswordForm account.UpdatePasswordForm

	if err:= c.ShouldBindJSON(&UpdatePasswordForm); err != nil{
		c.JSON(400, gin.H{"success":false})
		return
	}
	success := account.UpdatePassword(UpdatePasswordForm)
	if success{
		c.JSON(200, gin.H{"success":true, "msg":"Successfully update password"})
		return
	}
	c.JSON(200, gin.H{"success":true, "data":UpdatePasswordForm})
}*/

/*
func Test(c *gin.Context) {
	c.JSON(200, gin.H{"success": true, "result": "This is test"})
}*/
