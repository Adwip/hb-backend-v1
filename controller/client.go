package controller

import "github.com/gin-gonic/gin"
import "hb-backend-v1/model"

type clientObject struct {
}

func Client() *clientObject {
	clientObj := &clientObject{}

	return clientObj
}

func (clientObject) TestHandler(c *gin.Context) {
	c.JSON(200, model.WebResponse{Success: true, Msg: "Test function"})
}

func (clientObject) TestHandlerParams(c *gin.Context) {
	params := c.Param("id")
	c.JSON(200, model.WebResponse{Success: true, Data: params, Msg: "Test params"})
}
