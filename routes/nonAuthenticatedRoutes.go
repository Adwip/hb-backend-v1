package routes

import "github.com/gin-gonic/gin"
import "hb-backend-v1/controller"

func nonAuthenticatedRoutesRoutes(router *gin.Engine) {

	accountCtrl := controller.Account()
	auth := router.Group("/auth")
	{
		auth.POST("/login", accountCtrl.Login)
		auth.POST("/registration", accountCtrl.Regristration)
		auth.PUT("/password", accountCtrl.UpdatePassword)
		// auth.POST("/destroy",nil)
	}

	productCtrl := controller.Product()
	product := router.Group("/user-product")
	{
		product.POST("/", productCtrl.AddProduct)
	}
}
