package main

import "hb-backend-v1/routes"
import "hb-backend-v1/config"
import "fmt"
import "github.com/joho/godotenv"

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env file")
	}
	router := routes.Routes()
	database := config.Database()
	database.InitConnection()
	// config.InitDB()
	router.Run(":3001")
}
