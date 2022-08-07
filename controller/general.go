package controller

import "github.com/gin-gonic/gin"

type General struct {
}

func GeneralHandler() *General {
	handler := &General{}
	return handler
}

func (General) Options(c *gin.Context) {
	c.JSON(200, gin.H{"success": true})
}

func (General) NoRoute(c *gin.Context) {
	c.JSON(404, gin.H{"success": false, "msg": "URL Not Found"})
}
