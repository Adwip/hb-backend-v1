package routes

import (
	"hb-backend-v1/provider"

	"github.com/gin-gonic/gin"
)

func authenticatedRoutes(router *gin.Engine, handler *provider.HandlerInit, middleware *provider.MiddlewareInit) {

	// loginValidator := middleware.Login()
	authenticatedRoutes := router.Group("/ar", middleware.Auth.Logger)

	auth := authenticatedRoutes.Group("/auth")
	auth.PUT("/password", handler.Authentication.UpdatePassword)

	product := authenticatedRoutes.Group("/product")
	product.POST("/", handler.Product.AddProduct)
	/*
		productCtrl := controller.Product()
		product := authenticatedRoutes.Group("/user-product")
		{
			product.POST("/", productCtrl.AddProduct)

			product.POST("/add-design", nil)
			product.GET("/", nil)
			product.GET("/{ID}", nil)
			product.DELETE("/delete-list", nil)
			product.DELETE("/design", nil)
			product.PUT("/archive", nil)
			product.PUT("/keep", nil) //[on process discussion with client]
			product.PUT("/design", nil)
			product.GET("/image{ID}", nil)
			product.GET("/search", nil)
			product.GET("/list", nil)
			product.GET("/recommendation", nil)
		}
		user := authenticatedRoutes.Group("/user")
		{
			_ = user
			// user.POST("/all-user", account.AllAccount)
			// user.GET("/")
			// user.PUT("/update-profile")
			// user.GET("/{username}")
			// user.PUT("/archive")
			// user.GET("/rating")
			// user.GET("/review")
		}

		client := authenticatedRoutes.Group("/client")
		{
			client.POST("/profile")
			client.GET("/certtificate")
			client.GET("/product-bought")
			client.PUT("/basic")
			client.GET("/basic")
		}
			discussion := authenticatedRoutes.Group("/discussion")
			{
				discussion.POST("/")
				discussion.PUT("/archive")
				discussion.POST("/chat-history")
			}
	*/
	/*
		transaction := authenticatedRoutes.Group("/transaction")
		{
			transaction.POST("/")
			transaction.GET("/success")
			transaction.GET("/waiting")
			transaction.GET("/canceled")
			transaction.GET("/certificate")
			transaction.GET("/{ID-Certificate}")
		}
	*/
	/*
			payment := authenticatedRoutes.Group("/payment")
			{
				payment.POST("")
			}

		account := authenticatedRoutes.Group("/account")
		{
			_ = account
			// account.GET("/", controller.AllAccount)
		}
	}*/
}
