package main

import "hb-backend-v1/routes"
import "hb-backend-v1/config/database"


func main(){
	router := routes.Routes()
	database.InitDB()
	
	router.Run(":3001")
}