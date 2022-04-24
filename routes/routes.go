package routes

import "github.com/gin-gonic/gin"
import "hb-backend-v1/middleware"

func Routes() *gin.Engine {
	router := gin.New()
	corsValidator := middleware.CORS()

	// router.Use(loginRoute)
	router.Use(corsValidator.Cors)
	authenticatedRoutes(router)
	nonAuthenticatedRoutesRoutes(router)
	// router.Use(loginValidator.LoginChecking)

	router.OPTIONS("", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": 200})
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"success": false, "error": "URL Not Found"})
	})

	return router
}
