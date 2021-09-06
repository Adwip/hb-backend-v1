package main

import "hb-backend-v1/routes"


func main(){
	router := routes.Routes()

	router.Run()
}