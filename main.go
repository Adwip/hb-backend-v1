package main

import "hb-backend-v1/routes"
import "hb-backend-v1/config"
import "github.com/joho/godotenv"
import "github.com/gin-gonic/gin"
import "os"
import "hb-backend-v1/provider"

func main() {
	err := godotenv.Load()
	app := gin.New()

	if err != nil {
		panic("Failed to load .env file")
	}
	database := config.Database(os.Getenv("MY_SQL_URL"))
	//Init MySQL connection
	mySQL := database.InitMySQL()
	//Init Repos
	repoInit := provider.InitRepositories(mySQL)
	//Init Services
	serviceInit := provider.InitServices(repoInit)
	// Init middleware
	middleware := provider.InitMiddleware(repoInit)
	// Init Handlers
	handler := provider.InitHandlers(serviceInit)

	routes.Init(app, handler, middleware)
	// router.Run(":3001")
	app.Run(":3001")
}
