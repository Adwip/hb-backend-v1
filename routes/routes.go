package routes

import "github.com/gin-gonic/gin"
import "hb-backend-v1/middleware"
import "hb-backend-v1/controller"

func Routes() *gin.Engine {
	router := gin.New()
	corsValidator := middleware.CORS()
	handler := controller.Handler()

	// router.Use(loginRoute)
	router.Use(corsValidator.Cors)
	authenticatedRoutes(router)
	nonAuthenticatedRoutes(router)
	// router.Use(loginValidator.LoginChecking)

	router.OPTIONS("", handler.Options)
	router.NoRoute(handler.NoRoute)

	return router
}
