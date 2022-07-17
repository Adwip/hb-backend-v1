package routes

import "github.com/gin-gonic/gin"
import "hb-backend-v1/controller"

func nonAuthenticatedRoutes(router *gin.Engine) {

	accountCtrl := controller.Account()
	auth := router.Group("/auth")
	{
		auth.POST("/login", accountCtrl.Login)
		auth.POST("/registration", accountCtrl.Regristration)
		// auth.POST("/destroy",nil)
	}

	productCtrl := controller.Product()
	product := router.Group("/product")
	{
		product.GET("/recommendation", productCtrl.Recommendation)
		// product.GET("/:id", productCtrl.TestHandlerParams)
		// product.GET("/card/:id", productCtrl.TestHandlerParams)
		// product.GET("/favorites-total/:id", productCtrl.TestHandler)
		// product.GET("/carts-total/:id", productCtrl.TestHandler)
	}

	/*
		userCtrl := controller.User()
		user := router.Group("/user")
		{
			user.GET("/:id", userCtrl.TestHandler)
			user.GET("/card/:id", userCtrl.TestHandler)
			user.GET("/rating/:id", userCtrl.TestHandler)
			user.GET("/review/:id", userCtrl.TestHandler)
			user.GET("/product/:id", userCtrl.TestHandler)
			user.GET("/trans/:id", userCtrl.TestHandler)
			user.GET("/visited-client/:id", userCtrl.TestHandler)
		}*/
	/*
		clientCtrl := controller.Client()
		client := router.Group("/client")
		{
			client.GET("/", clientCtrl.TestHandler)
			client.GET("/card/:id", clientCtrl.TestHandler)
		}*/
}
