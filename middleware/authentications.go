package middleware

import (
	"fmt"
	"hb-backend-v1/model"
	"hb-backend-v1/repository"
	"hb-backend-v1/utils"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthenticationInt interface {
	Logger(*gin.Context)
	AccessChecking() bool
}

type authentication struct {
	redis repository.Account //Actually for redis, will replcae later
}

func AuthMiddleware(redis *repository.Account) AuthenticationInt {
	return &authentication{
		redis: *redis,
	}
}

func (authentication) Logger(c *gin.Context) {
	reqHeader := c.Request.Header
	JWT := utils.JsonWT()
	token, isset := reqHeader["Authorization"]
	jwtKey := os.Getenv("JWT_SECRET_KEY")

	if !isset {
		c.AbortWithStatusJSON(401, model.WebResponse{Success: false, Msg: "Access rejected 0"})
		return
	}

	//"Bearer" string exists or not
	subToken := strings.Split(token[0], " ")
	if len(subToken) > 1 {
		token[0] = subToken[1]
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

func (authentication) AccessChecking() bool {
	return true
}
