package routes

import "github.com/gin-gonic/gin"
import "hb-backend-v1/controller"

func Routes() *gin.Engine {
	router := gin.New()
	
	account := router.Group("/account")
	{
		account.GET("/",controller.AllAccount)
	}

	return router
}