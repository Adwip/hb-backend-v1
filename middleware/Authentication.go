package middleware

import "github.com/gin-gonic/gin"
import _ "hb-backend-v1/library/authentication"
import "hb-backend-v1/model"
import "hb-backend-v1/library"
import _ "context"
import "strings"
import "fmt"

type LoginMdw struct {
}

func Login() *LoginMdw {
	login := &LoginMdw{}
	return login
}

func (LoginMdw) LoginChecking(c *gin.Context) {
	reqHeader := c.Request.Header
	fullPath := c.FullPath()
	JWT := library.JsonWT()
	token, isset := reqHeader["Authorization"]

	if fullPath == "/auth/login" || fullPath == "/auth/registration" {
		c.Next()
		return
	}

	if !isset {
		c.AbortWithStatusJSON(401, model.WebResponse{Success: false, Msg: "Access rejected 0"})
		return
	}
	splittedToken := strings.Split(token[0], ".")
	if length := len(splittedToken); length != 3 {
		c.AbortWithStatusJSON(401, model.WebResponse{Success: false, Msg: "Access rejected 1"})
		return
	}

	header, payload, err := JWT.DecodeToken(token[0])
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(401, model.WebResponse{Success: false, Msg: "Access rejected 2"})
		return
	}
	isAuthorized := JWT.VerifiyToken(splittedToken[0], splittedToken[1], splittedToken[2], header)

	if isAuthorized {
		c.Set("JWTHeader", header)
		c.Set("JWTPayload", payload)
		c.Next()
		return
	}
	c.AbortWithStatusJSON(401, model.WebResponse{Success: false, Msg: "Access rejected 3"})
}

func (LoginMdw) AccessChecking() bool {
	return true
}
