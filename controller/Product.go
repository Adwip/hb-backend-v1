package controller

import "github.com/gin-gonic/gin"
import "hb-backend-v1/library"
import "hb-backend-v1/model"
import "hb-backend-v1/model/product"

type productObj struct {
}

func Product() *productObj {
	productObject := &productObj{}
	return productObject
}

func (productObj) AddProduct(c *gin.Context) {
	var reqBody product.AddProduct
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(200, model.WebResponse{Success: false})
		return
	}
	c.JSON(200, model.WebResponse{Success: true, Data: reqBody})
}

func (productObj) AddList(c *gin.Context) {

	c.JSON(200, gin.H{"success": true, "result": "Berhasil"})
}

func (productObj) TestTime(c *gin.Context) {
	identity := library.Identity(c)
	time := library.Time()
	timeConvert, errConvert := time.ConvertToUTC("20-12-2020 13:00:10", identity.GetTimezone())
	if errConvert != nil {
		c.JSON(200, model.WebResponse{Success: false, Msg: errConvert.Error()})
		return
	}
	c.JSON(200, model.WebResponse{Success: true, Data: timeConvert})
}
