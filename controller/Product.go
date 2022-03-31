package controller

import "github.com/gin-gonic/gin"

type product struct {
}

func Product() *product {
	productObject := &product{}
	return productObject
}

func (product) AddList(c *gin.Context) {

	c.JSON(200, gin.H{"success": true, "result": "Berhasil"})
}
