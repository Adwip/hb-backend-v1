package model

import "github.com/gin-gonic/gin"

type WebResponse struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
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
