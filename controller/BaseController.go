package controller

import "github.com/gin-gonic/gin"

type BaseController struct{
	c *gin.Context
}

/*
func (bc *BaseController) JSONParser () {

}*/

func (bc *BaseController) Response(success bool, payload interface{}){
	if success {
		bc.c.JSON(200, gin.H{"success":true, "data":payload})
		return
	}
	bc.c.JSON(500, gin.H{"success":true})
}
