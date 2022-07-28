package provider

type middlewareInit struct {
}

func InitMiddleware() *middlewareInit {
	return &middlewareInit{}
}
