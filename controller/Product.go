package controller

import "github.com/gin-gonic/gin"

func AddList(c *gin.Context){
	
	c.JSON(200, gin.H{"success":true, "result":"Berhasil"})
}