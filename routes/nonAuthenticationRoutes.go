package routes

import "github.com/gin-gonic/gin"
import "hb-backend-v1/controller"
import _ "hb-backend-v1/middleware"

func nonAuthenticatedRoutesRoutes(router *gin.Engine) {

	productCtrl := controller.Product()
	product := router.Group("/user-product")
	{
		product.POST("/", productCtrl.AddProduct)
	}
}
