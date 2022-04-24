package middleware

import "github.com/gin-gonic/gin"
import "hb-backend-v1/model"
import "hb-backend-v1/library"
import "strings"
import "fmt"
import "os"

type LoginMdw struct {
}

func Login() *LoginMdw {
	login := &LoginMdw{}
	return login
}

func (LoginMdw) Logger(c *gin.Context) {
	reqHeader := c.Request.Header
	JWT := library.JsonWT()
	token, isset := reqHeader["Authorization"]
	jwtKey := os.Getenv("JWT_SECRET_KEY")

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
	isAuthorized := JWT.VerifiyToken(splittedToken[0], splittedToken[1], splittedToken[2], header, jwtKey)

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
