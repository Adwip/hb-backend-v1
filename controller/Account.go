package controller

import "github.com/gin-gonic/gin"
import "hb-backend-v1/model/account"
import "hb-backend-v1/form/accountForm"

func AllAccount(c *gin.Context) {
	result, err := account.AllAccount()

	if err != nil {
		c.JSON(200, gin.H{"hasil": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true, "result": result})
	}
}

func Login(c *gin.Context) {

	var LoginForm accountForm.LoginForm
	// result, err := account.Login()
	if err := c.ShouldBindJSON(&LoginForm); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	exists, result, _ := account.Login(LoginForm.UnameMail, LoginForm.Password)
	if !exists {
		c.JSON(200, gin.H{"success": false, "data": result})
		return
	}
	c.JSON(200, gin.H{"success": true, "data": result})
}

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
}

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

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"success": true, "result": "This is test"})
}
