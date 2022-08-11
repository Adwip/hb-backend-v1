package provider

import "hb-backend-v1/service"

// import "fmt"

type serviceInit struct {
	Account service.Account
	Product service.Product
}

func InitServices(repoInit *repositoryInit) *serviceInit {
	return &serviceInit{
		Account: service.NewAccountService(&repoInit.Account),
		Product: service.NewProductService(&repoInit.Product, &repoInit.ProductImage),
	}
}
