package repository

/*
import "hb-backend-v1/config"
import "github.com/gin-gonic/gin"
import "database/sql"
import "hb-backend-v1/model/product"
import "github.com/google/uuid"
import "hb-backend-v1/model"
import "context"
import "time"
import _ "hb-backend-v1/library"
import "fmt"

type productFileRepo struct {
	conn *sql.DB
}

func ProductFile() *productFileRepo {
	db := config.Database()
	connection := db.GetConnection()
	productFile := &productFileRepo{
		conn: connection,
	}
	return productFile
}

func (pi productImageRepo) AddProductFiles(c *gin.Context, productID string, req []product.ProductImage) *model.RepoResponse {
	var rejectedFile []string
	var inserted int64
	var statement string
	var resultQuery sql.Result
	var errQuery, errRows error

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	for i := 0; i < len(req); i++ {
		statement = "INSERT INTO product_file (id_product_img, product, file_name, base_64) VALUES (?, ?, ?, ?)"
		resultQuery, errQuery = pi.conn.ExecContext(ctx, statement, uuid.New(), productID, req[i].ImageName, req[i].Base64)

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
		return &model.RepoResponse{Success: false, Data: rejectedFile}
	} else if len(rejectedFile) != 0 {
		return &model.RepoResponse{Success: true, Data: rejectedFile}
	}

	return &model.RepoResponse{Success: true}

}
*/
