package main

import "hb-backend-v1/routes"
import "hb-backend-v1/config"
import _ "fmt"
import "github.com/joho/godotenv"

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}
	router := routes.Routes()
	database := config.Database()
	database.InitConnection()
	router.Run(":3001")
}
