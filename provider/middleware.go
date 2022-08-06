package provider

import "hb-backend-v1/middleware"

type MiddlewareInit struct {
	Logger middleware.AuthenticationInt
}

func InitMiddleware(repos *repositoryInit) *MiddlewareInit {
	return &MiddlewareInit{
		Logger: middleware.AuthMiddleware(&repos.Account),
	}
}
