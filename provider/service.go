package provider

import "hb-backend-v1/service"

// import "fmt"

type serviceInit struct {
	Account service.Account
}

func InitServices(repoInit *repositoryInit) *serviceInit {
	return &serviceInit{
		Account: service.NewAccountService(&repoInit.Account),
	}
}
