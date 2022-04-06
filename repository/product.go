package repository

import "hb-backend-v1/config"
import "github.com/gin-gonic/gin"
import "database/sql"
import "hb-backend-v1/model"

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

func (pr productRepo) AddProduct(c *gin.Context) *model.RepoResponse {
	return &model.RepoResponse{Success: false}
}
