package Controller

import "github.com/gin-gonic/gin"
import "net/http"
import "hb-backend-v1/form"

type User struct{
	ID int
	Name string
	Status bool
}

func Login(c *gin.Context){

}

func signup(){

}

func destroy_session(){

}

func change_password(){
	
}

func UserInfo() User {
	var pengguna = User{2, "Ini User", true}
	return  pengguna
}