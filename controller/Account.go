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