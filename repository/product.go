package repository

import "hb-backend-v1/config"
import "github.com/gin-gonic/gin"
import "database/sql"
import "hb-backend-v1/model"
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
	connSring := database.GetConnection()
	product := &productRepo{
		conn: connSring,
	}
	return product
}

func (pr productRepo) AddProduct(c *gin.Context, req product.AddProduct) *model.RepoResponse {
	var negotiate int8
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	identity := library.Identity(c)
	currentTime := library.Time().CurrentDateTimeDbFormat()

	defer cancel()

	id := uuid.New()
	if req.Negotiate {
		negotiate = 1
	}
	fmt.Println(id)
	fmt.Println(identity.GetUserID())
	statement := "INSERT INTO product (id_product, user, field, judul, negosiasi, createdAt) VALUES (?, ?, ?, ?, ?, ?)"

	result, errInsert := pr.conn.ExecContext(ctx, statement, id, identity.GetUserID(), req.Field, req.Title, negotiate, currentTime)
	if errInsert != nil {
		fmt.Println(errInsert)
		return &model.RepoResponse{Success: false, Msg: "Failed to add product"}
	}

	inserted, errRows := result.RowsAffected()
	if errRows != nil {
		fmt.Println(errRows)
		return &model.RepoResponse{Success: false, Msg: "Failed to add product"}
	} else if inserted == 0 {
		fmt.Println("Inserted value ", inserted)
		return &model.RepoResponse{Success: false, Msg: "Failed to add product"}
	}
	return &model.RepoResponse{Success: true, Data: id}
}

func (pr productRepo) AddProductImage(c *gin.Context, imageRequest []product.ProductImage) *model.RepoResponse {
	var rejectedFile []string
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	id := uuid.New()
	defer cancel()

	for i := 0; i < len(imageRequest); i++ {
		statement := ""
		prepStatement, errPrepare := pr.conn.PrepareContext(ctx, statement)
		if errPrepare != nil {
			fmt.Println("Failed to prepare query")
		}
		result, err := prepStatement.ExecContext(ctx, id)

		if err != nil {
			fmt.Println("Failed to insert data")
		}
		affected, _ := result.RowsAffected()

		if affected == 0 {
			rejectedFile = append(rejectedFile, imageRequest[i].ImageName)
		}
	}

	if len(rejectedFile) != 0 {
		return &model.RepoResponse{Success: false, Data: rejectedFile}
	}

	return &model.RepoResponse{Success: true}
}
