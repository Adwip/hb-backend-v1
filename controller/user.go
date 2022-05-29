package controller

import "github.com/gin-gonic/gin"
import "hb-backend-v1/model"

type userObject struct {
}

func User() *userObject {
	userObj := &userObject{}

	return userObj
}

func (userObject) TestHandler(c *gin.Context) {
	c.JSON(200, model.WebResponse{Success: true, Msg: "Test function"})
}

func (userObject) TestHandlerParams(c *gin.Context) {
	params := c.Param("id")
	c.JSON(200, model.WebResponse{Success: true, Data: params, Msg: "Test params"})
}
