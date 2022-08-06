package provider

import "hb-backend-v1/service"

// import "fmt"

type serviceInit struct {
	Authentication service.AuthenticationInt
}

func InitServices(repoInit *repositoryInit) *serviceInit {
	return &serviceInit{
		Authentication: service.NewAuthentication(&repoInit.Account),
	}
}
