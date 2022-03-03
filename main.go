package main

import "hb-backend-v1/routes"
import "hb-backend-v1/config/database"
import _"fmt"


func main(){
	router := routes.Routes()
	database.InitDB()
	router.Run(":3001")
}