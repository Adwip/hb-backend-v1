package provider

import "github.com/gin-gonic/gin"

func InitRoutes(handler *handlerInit, router *gin.Engine) {
	handler.Authentication.Routes(router)
}
