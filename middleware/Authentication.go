package middleware

import "github.com/gin-gonic/gin"
import _ "hb-backend-v1/library/authentication"
import "hb-backend-v1/model"
import "hb-backend-v1/library"

type LoginMdw struct {
}

func Login() *LoginMdw {
	login := &LoginMdw{}
	return login
}

func (LoginMdw) LoginChecking(c *gin.Context) {
	var isAuthorized bool
	reqHeader := c.Request.Header
	fullPath := c.FullPath()
	JWT := library.JsonWT()
	identity := library.Identity()
	var token, isset = reqHeader["Authorization"]
	if fullPath == "/auth/login" || fullPath == "/auth/registration" {
		c.Next()
		return
	}

	if !isset {
		c.AbortWithStatusJSON(401, model.WebResponse{Success: false, Msg: "Access rejected"})
		return
	}

	header, payload, err := JWT.DecodeToken(token[0])
	if err != nil {
		c.AbortWithStatusJSON(401, model.WebResponse{Success: false, Msg: "Access rejected"})
		return
	}
	isAuthorized = JWT.VerifiyToken(token[0], token[1], token[2], header)

	if isAuthorized {
		identity.SetHeader(&header)
		identity.SetPayload(&payload)
		c.Next()
		return
	}
	c.AbortWithStatusJSON(401, model.WebResponse{Success: false, Msg: "Access rejected"})
}

func (LoginMdw) AccessChecking() bool {
	return true
}
