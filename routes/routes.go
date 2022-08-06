package routes

import "github.com/gin-gonic/gin"
import "hb-backend-v1/provider"

func Init(app *gin.Engine, handler *provider.HandlerInit, middleware *provider.MiddlewareInit) {

	// corsValidator := middleware.CORS()
	// handler := controller.Handler()
	authenticatedRoutes(app, handler, middleware)
	nonAuthenticatedRoutes(app, handler)

	// router.Use(loginRoute)
	// router.Use(corsValidator.Cors)
	// router.Use(loginValidator.LoginChecking)

	// router.OPTIONS("", handler.Options)
	// router.NoRoute(handler.NoRoute)

}
