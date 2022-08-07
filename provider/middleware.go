package provider

import "hb-backend-v1/middleware"

type MiddlewareInit struct {
	Auth middleware.AuthenticationInt
}

func InitMiddleware(repos *repositoryInit) *MiddlewareInit {
	return &MiddlewareInit{
		Auth: middleware.AuthMiddleware(&repos.Account),
	}
}
