package controller

import "github.com/gin-gonic/gin"
import "hb-backend-v1/model"
import "hb-backend-v1/service"

// import "fmt"
// import "reflect"

type ProductController struct {
	product service.Product
}

func NewProductController(product *service.Product) *ProductController {
	return &ProductController{
		product: *product,
	}
}

func (handler ProductController) AddProduct(c *gin.Context) {
	var form model.AddProductRequest
	err := c.BindJSON(&form)
	if err != nil {
		c.JSON(200, model.WebResponse{Success: false})
		return
	}

	success, result, msg := handler.product.AddProduct(c, &form)

	if !success {
		c.JSON(200, model.WebResponse{Success: false, Msg: msg})
		return
	}
	// fmt.Println(reflect.TypeOf(insert.Data))

	_, imagesFail := handler.product.AddProductImages(c, &form.Images, result)

	c.JSON(200, model.WebResponse{Success: true, Data: imagesFail, Msg: msg})
}

/*
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
*/
