package main

// import "fmt"
// import "net/http"
// import auth "hb-backend-v1/controller/Authentication"
// import user "hb-backend-v1/controller/User"
// import "encoding/json"
// import "hb-backend-v1/routes"
// import "github.com/gin-gonic/gin"
// import "reflect"
import "hb-backend-v1/routes"

/*
type Response struct {
	auth.Auth
	user.User
}*/

// func handleReq(w http.ResponseWriter, r *http.Request){
	// var hasil []byte
	// w.Header().Set("Content-Type", "application/json")
	// var resp = Response{auth.JsonTest(), user.UserInfo()}
	// auth := auth.JsonTest()
	// user := user.UserInfo()
	// fmt.Fprintln(w, "Ini controller handleReq")
	// JsonResult, _ := json.Marshal(user.UserInfo())
	// w.Write(JsonResult)
	/*
	if err !=nil {
		w.Write(JsonResult)
	}else{
		fmt.Fprintln(w, err.Error())
	}*/
// }

func main(){
	// fmt.Println("Hello world ")
	// http.HandleFunc("/",handleReq)

	// fmt.Println("starting web server at http://localhost:8080/")
	// http.ListenAndServe(":8080", nil)
	// routes.Run(":8080")
	// router := routes.Routes()
	// router.Run(":8080")
	router := routes.Routes()

	router.Run()
	// fmt.Println(reflect.TypeOf(gin.Default()))
}