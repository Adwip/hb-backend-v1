package service

import (
	"hb-backend-v1/model"
	"hb-backend-v1/repository"

	"github.com/gin-gonic/gin"
)

type Product interface {
	AddProduct(*gin.Context, *model.AddProductRequest) (bool, string, string)
	AddProductImages(*gin.Context, *[]model.ProductImage, string) (bool, []string)
	ProductRecommended(*gin.Context) []model.AllProductsResponse
	// UpdateProduct()
	// DeleteProduct()
}

type ProductService struct {
	productRepo      repository.Product
	productImageRepo repository.ProductImage
}

func NewProductService(product *repository.Product, productImage *repository.ProductImage) Product {
	return &ProductService{
		productRepo:      *product,
		productImageRepo: *productImage,
	}
}

func (service ProductService) AddProduct(c *gin.Context, req *model.AddProductRequest) (bool, string, string) {
	success, id, msg := service.productRepo.AddProduct(c, req)
	return success, id, msg
}

func (service ProductService) AddProductImages(c *gin.Context, req *[]model.ProductImage, productID string) (bool, []string) {

	return service.productImageRepo.AddProductImages(c, productID, *req)
}

func (service ProductService) ProductRecommended(c *gin.Context) []model.AllProductsResponse {
	return service.productRepo.ProductRecommended(c)
}
