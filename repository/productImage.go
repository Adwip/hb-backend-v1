package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hb-backend-v1/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductImage interface {
	AddProductImages(*gin.Context, string, []model.ProductImage) (bool, []string)
}

type ProductImageRepo struct {
	db *sql.DB
}

func NewProductImageRepo(db *sql.DB) ProductImage {
	return &ProductImageRepo{
		db: db,
	}
}

func (repo ProductImageRepo) AddProductImages(c *gin.Context, productID string, req []model.ProductImage) (bool, []string) {
	var rejectedFile []string
	var inserted int64
	var statement string
	var resultQuery sql.Result
	var errQuery, errRows error

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	for i := 0; i < len(req); i++ {
		statement = "INSERT INTO product_image (id_product_img, product, file_name, base_64) VALUES (?, ?, ?, ?)"
		resultQuery, errQuery = repo.db.ExecContext(ctx, statement, uuid.New(), productID, req[i].ImageName, req[i].Base64)

		if errQuery != nil {
			fmt.Println(errQuery)
		}
		inserted, errRows = resultQuery.RowsAffected()

		if errRows != nil {
			fmt.Println(errRows)
		}

		if inserted == 0 {
			rejectedFile = append(rejectedFile, req[i].ImageName)
		}
	}

	if len(rejectedFile) != 0 && len(rejectedFile) == len(req) {
		return false, rejectedFile
	}

	return true, rejectedFile

}

/*
func (pi productImageRepo) ImagesByProductID(c *gin.Context, id string) (bool, []model.ProductImagesResponse) {
	var result []model.ProductImagesResponse

	sqlStatement := `SELECT * from product_image where product = ?`

	_ = sqlStatement

	return true, result
}*/
