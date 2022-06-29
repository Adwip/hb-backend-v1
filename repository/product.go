package repository

import "hb-backend-v1/config"
import "github.com/gin-gonic/gin"
import "database/sql"
import "context"
import "time"
import "github.com/google/uuid"
import "hb-backend-v1/model/product"
import "hb-backend-v1/library"
import "fmt"

type productRepo struct {
	conn *sql.DB
}

func Product() *productRepo {
	database := config.Database()
	product := &productRepo{
		conn: database.GetConnection(),
	}
	return product
}

func (pr productRepo) AddProduct(c *gin.Context, req product.AddProduct) (bool, string, string) {
	var negotiate int8
	var ImageError, purchaseError error
	var failedImages int
	var productImage product.ProductImage
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	identity := library.Identity(c)
	currentTime := library.Time().CurrentDateTimeDbFormat()

	defer cancel()

	productID := uuid.New()
	if req.OpenNegotiate {
		negotiate = 1
	}

	tx, err := pr.conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		fmt.Println(err)
		return false, "", "Failed to add product"
	}

	productStat := "INSERT INTO product (id_product, user, field, title, negotiate, purchaseType, type, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, ErrProduct := tx.Exec(productStat, productID, identity.GetUserID(), req.Field, req.Title, negotiate, req.PurchaseType, req.Type, currentTime)
	if ErrProduct != nil {
		tx.Rollback()
		fmt.Println(ErrProduct)
	}

	if req.PurchaseType == "SPC" {
		SinglePurchaseStat := "INSERT INTO one_time_purchase (id_spc, product, harga, status) VALUES (?, ?, ?, ?)"
		_, purchaseError = tx.Exec(SinglePurchaseStat, uuid.New(), productID, req.Price, req.Status)
	} else {
		MultiPurchaseStat := "INSERT INTO multiple_purchase (id_mpc, product, kuota, harga, status) VALUES (?, ?, ?, ?, ?)"
		_, purchaseError = tx.Exec(MultiPurchaseStat, uuid.New(), productID, req.Kuota, req.Price, req.Status)
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

func (pr productRepo) ProductByID(c *gin.Context, id string) (bool, product.ProductByIdResponse) {
	var result product.ProductByIdResponse
	return true, result
}
