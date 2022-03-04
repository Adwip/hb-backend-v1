package main

import "hb-backend-v1/routes"
import "hb-backend-v1/config/database"
import "fmt"
import "github.com/joho/godotenv"


func main(){
	router := routes.Routes()
	err := godotenv.Load()
	if err!=nil{
		fmt.Println("Failed to load .env file")
	}
	database.InitDB()
	router.Run(":3001")
}