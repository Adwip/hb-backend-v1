package controller

import "github.com/gin-gonic/gin"
import "hb-backend-v1/library"
import "hb-backend-v1/model"
import "hb-backend-v1/model/product"
import "hb-backend-v1/repository"
import "fmt"
import "reflect"

type productObj struct {
}

func Product() *productObj {
	productObject := &productObj{}
	return productObject
}

func (productObj) AddProduct(c *gin.Context) {
	var reqBody product.AddProduct
	var failedInsert map[string]interface{} = map[string]interface{}{}
	// failedInsert = map[string]interface{{}
	producRepo := repository.Product()
	productImageRepo := repository.ProductImage()
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(200, model.WebResponse{Success: false})
		return
	}

	insert := producRepo.AddProduct(c, reqBody)

	if !insert.Success {
		c.JSON(200, model.WebResponse{Success: false, Msg: insert.Msg})
		return
	}
	fmt.Println(reflect.TypeOf(insert.Data))
	insertImage := productImageRepo.AddProductImages(c, insert.Data.(string), reqBody.Images)
	if insertImage.Data != nil {
		failedInsert["productImage"] = insertImage.Data
	}
	c.JSON(200, model.WebResponse{Success: true, Data: failedInsert})
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