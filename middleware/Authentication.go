package middleware

import "github.com/gin-gonic/gin"
import _ "hb-backend-v1/library/authentication"
import "hb-backend-v1/model"
import "hb-backend-v1/library"
import _ "context"
import "strings"
import "fmt"
import "os"

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
	jwtKey := os.Getenv("JWT_SECRET_KEY")

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

func (LoginMdw) CORS(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("Access-Control-Allow-Origin", os.Getenv("ALLOW_ORIGIN"))
	c.Header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Range,Content-Disposition, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()
}
