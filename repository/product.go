package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hb-backend-v1/model"
	"hb-backend-v1/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Product interface {
	AddProduct(*gin.Context, *model.AddProductRequest) (bool, string, string)
	ProductRecommended(*gin.Context) []model.AllProductsResponse
}

type ProductRepo struct {
	conn *sql.DB
}

func NewProductRepo(db *sql.DB) Product {
	return &ProductRepo{
		conn: db,
	}
}

func (repo ProductRepo) AddProduct(c *gin.Context, req *model.AddProductRequest) (bool, string, string) {
	var negotiate int8
	var ImageError, purchaseError error
	var failedImages int
	var productImage model.ProductImage
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	identity := utils.Identity(c)
	currentTime := utils.Time().CurrentDateTimeDbFormat()

	defer cancel()

	productID := uuid.New()
	if req.OpenNegotiate {
		negotiate = 1
	}

	tx, err := repo.conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		fmt.Println(err)
		return false, "", "Failed to add product"
	}

	productStat := "INSERT INTO product (id_product, user, field, title, negotiable, purchaseType, type, price, status, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, ErrProduct := tx.Exec(productStat, productID, identity.GetUserID(), req.Field, req.Title, negotiate, req.PurchaseType, req.Type, req.Price, req.Status, currentTime)
	if ErrProduct != nil {
		tx.Rollback()
		fmt.Println(ErrProduct)
	}

	if req.PurchaseType == "SPC" {
		SinglePurchaseStat := "INSERT INTO one_time_purchase (id_spc, product) VALUES (?, ?)"
		_, purchaseError = tx.Exec(SinglePurchaseStat, uuid.New(), productID)
	} else {
		MultiPurchaseStat := "INSERT INTO multiple_purchase (id_mpc, product, kuota) VALUES (?, ?, ?)"
		_, purchaseError = tx.Exec(MultiPurchaseStat, uuid.New(), productID, req.Kuota)
	}

	if purchaseError != nil {
		tx.Rollback()
		fmt.Println(purchaseError)
	}

	imageStat := "INSERT INTO product_image (id_product_img, product, file_name, base_64) VALUES (?, ?, ?, ?)"
	for i := 0; i < len(req.Images); i++ {
		productImage = req.Images[i]
		_, ImageError = tx.Exec(imageStat, uuid.New(), productID, productImage.ImageName, productImage.Base64)
		if ImageError != nil {
			fmt.Println(ImageError)
			failedImages++
		}
	}

	if failedImages == len(req.Images) {
		tx.Rollback()
	}

	errTrans := tx.Commit()

	if errTrans != nil {
		fmt.Println(errTrans)
		return false, "", "Failed to add product"

	}

	return true, productID.String(), ""
}

func (pr ProductRepo) ProductRecommended(c *gin.Context) []model.AllProductsResponse {

	var result []model.AllProductsResponse
	var row model.AllProductsResponse
	var errorRow error
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	sqlStatement := `
	SELECT
		id_product as id,
		title as productName,
		firstName as creator,
		negotiable,
		purchaseType as purchaseType,
		count(id_favourite) as favorite,
		price,
		IFNULL(product_image.file_name, "") as productImage

	FROM product
		INNER JOIN user ON user.id_user = product.user
		INNER JOIN account_information ON account_information.account = user.account
		LEFT JOIN product_image ON product_image.product = product.id_product
		LEFT JOIN favourite ON favourite.product = product.id_product
		LEFT JOIN main_image ON main_image.image = product_image.id_product_img
		LEFT JOIN one_time_purchase ON one_time_purchase.product = product.id_product
		LEFT JOIN multiple_purchase ON multiple_purchase.product = product.id_product
	WHERE
		(multiple_purchase.kuota IS NOT NULL OR multiple_purchase.kuota > 0) OR (product.purchaseType = 'SPC' AND one_time_purchase.offerStatus = 'On Offering')
	GROUP BY product.id_product
	`

	rows, err := pr.conn.QueryContext(ctx, sqlStatement)

	if err != nil {
		fmt.Println(err)
		return result
	}

	for rows.Next() {
		errorRow = rows.Scan(&row.ID, &row.ProductName, &row.Creator, &row.Negotiable, &row.PurchaseType, &row.Favourite, &row.Price, &row.ProductImage)
		if errorRow != nil {
			fmt.Println(errorRow)
			return result
		}

		result = append(result, row)
	}

	return result
}

/*
func (pr productRepo) DetailByID(c *gin.Context, id string) (bool, model.ProductByIDResponse) {

	var result model.ProductByIDResponse
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	sqlStatement := `
	SELECT
		id_product as id,
		title as productName,
		firstName as creator,
		negotiable,
		purchaseType as purchaseType,
		count(id_favourite) as favorite,
		price,
		IFNULL(offerStatus, "") as offerStatus,
		product.status,
		IFNULL(product_image.file_name, "") as productImage

	FROM product
		INNER JOIN user ON user.id_user = product.user
		INNER JOIN account_information ON account_information.account = user.account
		LEFT JOIN product_image ON product_image.product = product.id_product
		LEFT JOIN favourite ON favourite.product = product.id_product
		LEFT JOIN main_image ON main_image.image = product_image.id_product_img
		LEFT JOIN one_time_purchase ON one_time_purchase.product = product.id_product
		LEFT JOIN multiple_purchase ON multiple_purchase.product = product.id_product
	WHERE
		((multiple_purchase.kuota IS NOT NULL OR multiple_purchase.kuota > 0) OR (product.purchaseType = 'SPC' AND one_time_purchase.offerStatus = 'On Offering')) AND id_product = ?
	GROUP BY product.id_product
	`

	row := pr.conn.QueryRowContext(ctx, sqlStatement, id)

	err := row.Scan(&result.ID, &result.ProductName, &result.Creator, &result.Negotiable, &result.PurchaseType, &result.Favourite, &result.Price, &result.ProductImage, &result.OfferStatus, &result.Status)

	if err != nil {
		fmt.Println(err)
		return false, result
	}

	return true, result
}*/

/*
func (pr productRepo) ProductByID(c *gin.Context, id string) (bool, product.ProductByIdResponse) {
	var result product.ProductByIdResponse
	var row product.ProductByIdResponse

	ctx, cancel := context.WithTimeout(c, 5*time.Second)

	defer cancel()

	sqlStatement := `
	SELECT
		title as productName,
		firstName as creator,
		negotiable,
		count(id_favourite) as favorite,
		price,
		image as productImage

	FROM product
		INNER JOIN user ON user.id_user = product.user
		INNER JOIN account_information ON account_information.account = user.account
		LEFT JOIN product_image ON product_image.product = product.id_product
		LEFT JOIN favourite ON favourite.product = product.id_product
		LEFT JOIN main_image ON main_image.image = product_image.id_product_img
		LEFT JOIN one_time_purchase ON one_time_purchase.product = product.id_product
		LEFT JOIN multiple_purchase ON multiple_purchase.product = product.id_product
	WHERE
		(multiple_purchase.kuota IS NOT NULL OR multiple_purchase.kuota > 0) OR (product.purchaseType = 'SPC' AND one_time_purchase.offerStatus = 'On Offering')
	GROUP BY product.id_product
	`

	rows, err := pr.conn.QueryContext(c, sqlStatement)
	_ = rows
	if err!=nil {
		return false, result
	}

	for rows.Next() {
		row = rows.Scan()
	}


	return true, result
}*/
