package routes

import "github.com/gin-gonic/gin"
import "hb-backend-v1/controller"

func Routes() *gin.Engine {
	router := gin.New()

	// router.GET("/all-account", controller.AllAccount)
	// router.POST("/login",controller.Login)
	router.GET("/",controller.Test)
	
	
	auth := router.Group("/auth")
	{
		auth.POST("/login",controller.Login)
		// auth.POST("/registration",nil)
		// auth.PUT("/password",nil)
		// auth.POST("/destroy",nil)
	}
	
	/*
	product := router.Group("/product")
	{
		product.POST("/add-to-list",nil)
		product.POST("/add-design",nil)
		product.GET("/",nil)
		product.GET("/{ID}",nil)
		product.DELETE("/delete-list",nil)
		product.DELETE("/design",nil)
		product.PUT("/archive",nil)
		product.PUT("/keep",nil)//[on process discussion with client]
		product.PUT("/design",nil)
		product.GET("/image{ID}",nil)
		product.GET("/search",nil)
		product.GET("/list",nil)
		product.GET("/recommendation",nil)
	}
	*/
	/*
	user := router.Group("/user")
	{
		user.POST("/profile")
		user.GET("/")
		user.PUT("/update-profile")
		user.GET("/{username}")
		user.PUT("/archive")
		user.GET("/rating")
		user.GET("/review")
	}
	*/
	/*
	client := router.Group("/client")
	{
		client.POST("/profile")
		client.GET("/certtificate")
		client.GET("/product-bought")
		client.PUT("/basic")
		client.GET("/basic")
	}
	*/
	/*
	discussion := router.Group("/discussion")
	{
		discussion.POST("/")
		discussion.PUT("/archive")
		discussion.POST("/chat-history")
	}
	*/
	/*
	transaction := router.Group("/transaction")
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
	payment := router.Group("/payment")
	{
		payment.POST("")
	}
	*/

	account := router.Group("/account")
	{
		account.GET("/",controller.AllAccount)
	}

	return router
}