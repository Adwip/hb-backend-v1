package controller

import "github.com/gin-gonic/gin"
import "hb-backend-v1/model/account"

func AllAccount(c *gin.Context){
	result, err := account.AllAccount()
	
	if err!= nil{
		c.JSON(200, gin.H{"hasil":err.Error()})
	}else{
		c.JSON(200, gin.H{"success":true, "result":result})
	}
}


func Login(c*gin.Context){

	var LoginForm account.LoginForm
	// result, err := account.Login()
	if err := c.ShouldBindJSON(&LoginForm); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	exists, result, _ := account.Login(LoginForm.Username, LoginForm.Email, LoginForm.Password)
	if !exists {
		c.JSON(200, gin.H{"success":false})
		return
	}
	c.JSON(200, gin.H{"success":true, "data":result})
}

func Regristration(c *gin.Context){
	var RegistrationForm account.RegistrationForm

	if err := c.ShouldBindJSON(&RegistrationForm); err != nil{
		c.JSON(400, gin.H{"success":false, "message":err.Error()})
		return
	}
	success,err := account.RegistrationUser(RegistrationForm)
	if success{
		c.JSON(200, gin.H{"success":true, "data":RegistrationForm})
		return
	}
	c.JSON(400, gin.H{"success":false, "message":err.Error()})
}

func UpdatePassword(c *gin.Context){
	var UpdatePasswordForm account.UpdatePasswordForm

	if err:= c.ShouldBindJSON(&UpdatePasswordForm); err != nil{
		c.JSON(400, gin.H{"success":false})
		return
	}
	// success, err := account.UpdatePassword()
	c.JSON(200, gin.H{"success":true, "data":UpdatePasswordForm})
}

func Test(c *gin.Context){
	c.JSON(200, gin.H{"success":true, "result":"This is test"})
}