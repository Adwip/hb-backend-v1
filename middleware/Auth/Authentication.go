package auth

import "github.com/gin-gonic/gin"
import "fmt"

func LoginChecking(c *gin.Context){
	var isAuthorized bool
	var header = c.Request.Header
	var fullPath = c.FullPath()
	var token, isset = header["Authorization"]

	if isset{
		isAuthorized,_ = authentication.VerifyToken(token[0])
	}else{
		isAuthorized = false
	}
	if isAuthorized || fullPath == "/auth/login"{
		c.Next()
		return
	}
	c.AbortWithStatusJSON(401, gin.H{"error": "Akses ditolak"})
}

