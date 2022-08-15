package routes

import "github.com/gin-gonic/gin"
import "hb-backend-v1/provider"

func nonAuthenticatedRoutes(router *gin.Engine, handler *provider.HandlerInit) {

	authGroup := router.Group("/auth")
	authGroup.POST("/", handler.Authentication.Login)
	authGroup.POST("/registration", handler.Authentication.Regristration)
	// authGroup.PUT("/password", handler.Authentication.Regristration)
	// auth.POST("/destroy",nil)

	product := router.Group("/product")
	product.GET("/", handler.Product.Recommendation)
	/*
		productCtrl := controller.Product()
		product := router.Group("/product")
		{
			product.GET("/recommendation", productCtrl.Recommendation)
			product.GET("/:id", productCtrl.DetailByID)
			product.GET("/images/:id", productCtrl.DetailByID)
			// product.GET("/card/:id", productCtrl.TestHandlerParams)
			// product.GET("/favorites-total/:id", productCtrl.TestHandler)
			// product.GET("/carts-total/:id", productCtrl.TestHandler)
		}

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
