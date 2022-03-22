package main

import "hb-backend-v1/routes"
import "hb-backend-v1/config"
import "fmt"
import "github.com/joho/godotenv"

func main() {
	router := routes.Routes()
	err := godotenv.Load()
	database := config.Database()
	database.InitDB()
	if err != nil {
		fmt.Println("Failed to load .env file")
	}
	// config.InitDB()
	router.Run(":3001")
}
