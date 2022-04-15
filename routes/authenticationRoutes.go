package routes

import "github.com/gin-gonic/gin"
import "hb-backend-v1/controller"
import "hb-backend-v1/middleware"

func authenticatedRoutes(router *gin.Engine) {

	loginValidator := middleware.Login()
	loginRoutes := router.Group("/ar", loginValidator.LoginChecking)
	{
		productCtrl := controller.Product()
		product := loginRoutes.Group("/user-product")
		{
			product.POST("/", productCtrl.AddProduct)
		}
	}
}
