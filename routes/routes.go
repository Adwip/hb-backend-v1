package routes

import "github.com/gin-gonic/gin"
import "net/http"
import _"hb-backend-v1/controller"
import "fmt"
import "hb-backend-v1/config/database"

func Routes() *gin.Engine {
	router := gin.Default()
	router.GET("/",func(c *gin.Context){
		// c.String(http.StatusOK, fmt.Sprintf("Request GET success!"))
		hasil, err := sqlConnect.Query()
		fmt.Println(hasil)
		if err != nil{
			c.JSON(200, gin.H{"hasil":err.Error()})
		}else{
			c.JSON(200, gin.H{"success":true, "result":hasil})
		}
	})
	user := router.Group("/user")
	{
		user.POST("/signup",)
		user.POST("/login")
		user.POST("/change-password")
		user.GET("/info/:id")
	}
	v1 := router.Group("/v1")
	{
		v1.POST("/login",func(c *gin.Context){
			c.String(http.StatusOK, fmt.Sprintf("Request v1 login POST success!"))
		})
		v1.POST("/submit",func(c *gin.Context){
			c.String(http.StatusOK, fmt.Sprintf("Request v1 submit POST success!"))
		})
		v1.POST("/read",func(c *gin.Context){
			c.String(http.StatusOK, fmt.Sprintf("Request v1 read POST success!"))
		})
	}

	return router
}