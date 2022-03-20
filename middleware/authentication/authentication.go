package authentication

import "github.com/gin-gonic/gin"
import "hb-backend-v1/library/authentication"
import "hb-backend-v1/model"

func LoginChecking(c *gin.Context) {
	var isAuthorized bool
	var header = c.Request.Header
	var fullPath = c.FullPath()
	var token, isset = header["Authorization"]

	if isset {
		isAuthorized, _ = authentication.VerifyToken(token[0])
	} else {
		isAuthorized = false
	}
	if isAuthorized || fullPath == "/auth/login" || fullPath == "/auth/registration" {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(401, model.WebResponse{Success: false, Msg: "Access rejected"})
}
