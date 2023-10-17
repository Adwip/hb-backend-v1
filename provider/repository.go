package provider

import (
	"database/sql"
	"hb-backend-v1/repository"
)

type repositoryInit struct {
	Account      repository.Account
	Product      repository.Product
	ProductImage repository.ProductImage
}

func InitRepositories(db *sql.DB) *repositoryInit {
	return &repositoryInit{
		Account:      repository.NewAccountRepo(db),
		Product:      repository.NewProductRepo(db),
		ProductImage: repository.NewProductImageRepo(db),
	}
}
