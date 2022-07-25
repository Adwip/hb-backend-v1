package main

// import _ "hb-backend-v1/routes"
import "hb-backend-v1/config"
import "github.com/joho/godotenv"
import "github.com/gin-gonic/gin"
import "os"
import "hb-backend-v1/provider"

func main() {
	err := godotenv.Load()
	router := gin.New()

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
	// Init Handlers
	handler := provider.InitHandlers(serviceInit)
	// Init Routes
	provider.InitRoutes(handler, router)

	router.Run(":3001")
}
