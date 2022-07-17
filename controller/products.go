package controller

import "github.com/gin-gonic/gin"
import "hb-backend-v1/library"
import "hb-backend-v1/model"
import "hb-backend-v1/model/product"
import "hb-backend-v1/repository"

// import "fmt"
// import "reflect"

type productObj struct {
}

func Product() *productObj {
	productObject := &productObj{}
	return productObject
}

func (productObj) AddProduct(c *gin.Context) {
	var reqBody product.AddProduct
	producRepo := repository.Product()
	err := c.BindJSON(&reqBody)

	if err != nil {
		c.JSON(200, model.WebResponse{Success: false})
		return
	}

	success, id, msg := producRepo.AddProduct(c, reqBody)

	if !success {
		c.JSON(200, model.WebResponse{Success: false, Msg: msg})
		return
	}
	// fmt.Println(reflect.TypeOf(insert.Data))
	/*
		insertImage := productImageRepo.AddProductImages(c, id, reqBody.Images)
		if insertImage.Data != nil {
			failedInsert["productImage"] = insertImage.Data
		}*/
	c.JSON(200, model.WebResponse{Success: true, Data: id})
}

func (productObj) Recommendation(c *gin.Context) {
	productRepo := repository.Product()

	exist, result := productRepo.RecommendationProduct(c)

	if !exist {
		c.JSON(400, model.WebResponse{Success: false, Data: result})
		return
	}
	c.JSON(200, model.WebResponse{Success: true, Data: result})
}

func (productObj) DetailByID(c *gin.Context) {
	productRepo := repository.Product()

	exists, result := productRepo.DetailByID(c, c.Param("id"))
	if !exists {
		c.JSON(400, model.WebResponse{Success: false, Msg: "Product not found"})
		return
	}
	c.JSON(200, model.WebResponse{Success: true, Data: result})
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

func (productObj) TestHandler(c *gin.Context) {
	c.JSON(200, model.WebResponse{Success: true, Msg: "Test function"})

}

func (productObj) TestHandlerParams(c *gin.Context) {
	params := c.Param("id")
	c.JSON(200, model.WebResponse{Success: true, Data: params, Msg: "Test params"})
}
