package routes

import "hb-backend-v1/controller"
import "github.com/gin-gonic/gin"
import "hb-backend-v1/middleware"
import _ "hb-backend-v1/middleware/authentication"

func Routes() *gin.Engine {
	router := gin.New()

	loginValidator := middleware.Login()
	corsValidator := middleware.CORS()

	router.Use(corsValidator.Cors)
	router.Use(loginValidator.LoginChecking)

	accountCtrl := controller.Account()
	auth := router.Group("/auth")
	{
		auth.POST("/login", accountCtrl.Login)
		auth.POST("/registration", accountCtrl.Regristration)
		auth.PUT("/password", accountCtrl.UpdatePassword)
		// auth.POST("/destroy",nil)
	}

	productCtrl := controller.Product()
	product := router.Group("/product")
	{
		product.POST("/add-to-list", productCtrl.AddProduct)
		// product.GET("/add-to-list", productCtrl.TestTime)
		/*
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
			product.GET("/recommendation",nil)*/
	}

	user := router.Group("/user")
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
		_ = account
		// account.GET("/", controller.AllAccount)
	}

	router.OPTIONS("", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": 200})
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"success": false, "error": "URL Not Found"})
	})

	return router
}
