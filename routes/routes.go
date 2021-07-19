package routes

import "github.com/gin-gonic/gin"
import "net/http"

func routes(){
	router := gin.Default()

	auth := router.Group("/auth"){
		auth.POST("/login")
	}
}