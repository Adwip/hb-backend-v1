package provider

import handler "hb-backend-v1/controller"

type HandlerInit struct {
	Authentication handler.AccountCtrl
	General        handler.General
	Product        handler.ProductController
}

func InitHandlers(service *serviceInit) *HandlerInit {
	return &HandlerInit{
		Authentication: *handler.Account(&service.Account),
		General:        *handler.GeneralHandler(),
		Product:        *handler.NewProductController(&service.Product),
	}
}
