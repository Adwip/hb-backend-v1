package routes

import "github.com/gin-gonic/gin"
import _ "hb-backend-v1/middleware"
import _ "hb-backend-v1/controller"

func Routes() *gin.Engine {
	router := gin.New()
	authenticatedRoutes(router)
	nonAuthenticatedRoutes(router)
	/*
		corsValidator := middleware.CORS()
		handler := controller.Handler()

		// router.Use(loginRoute)
		router.Use(corsValidator.Cors)
		// router.Use(loginValidator.LoginChecking)

		router.OPTIONS("", handler.Options)
		router.NoRoute(handler.NoRoute)
	*/
	return router
}
