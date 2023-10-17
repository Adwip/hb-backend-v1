package routes

import (
	"hb-backend-v1/provider"

	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine, handler *provider.HandlerInit, middleware *provider.MiddlewareInit) {

	// corsValidator := middleware.CORS()
	// handler := controller.Handler()
	authenticatedRoutes(app, handler, middleware)
	nonAuthenticatedRoutes(app, handler)

	// router.Use(loginRoute)
	// router.Use(corsValidator.Cors)
	// router.Use(loginValidator.LoginChecking)

	app.OPTIONS("", handler.General.Options)
	app.NoRoute(handler.General.NoRoute)

}
