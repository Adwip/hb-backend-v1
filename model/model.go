package model

import "github.com/gin-gonic/gin"

type response int

const (
	SUCCESS response = 200
	FAILED  response = 500
)

type WebResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
}

type RepoResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
}

func Success(c *gin.Context, data interface{}, msg string) {
	result := WebResponse{
		Success: true,
		Msg:     msg,
		Data:    data,
	}
	c.JSON(200, result)
}

func Failed(c *gin.Context, data interface{}, msg string) {
	result := WebResponse{
		Success: true,
		Msg:     msg,
		Data:    data,
	}
	c.JSON(200, result)
}
