package controller

import "github.com/gin-gonic/gin"
import "hb-backend-v1/library"
import "hb-backend-v1/model"
import "hb-backend-v1/model/product"
import "hb-backend-v1/repository"

type productObj struct {
}

func Product() *productObj {
	productObject := &productObj{}
	return productObject
}

func (productObj) AddProduct(c *gin.Context) {
	var reqBody product.AddProduct
	producModel := repository.Product()
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(200, model.WebResponse{Success: false})
		return
	}

	insert := producModel.AddProduct(c, reqBody)

	if !insert.Success {
		c.JSON(200, model.WebResponse{Success: false, Msg: insert.Msg})
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
