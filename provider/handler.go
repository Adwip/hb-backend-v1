package provider

import handler "hb-backend-v1/controller"

type HandlerInit struct {
	Authentication handler.AccountCtrl
	General        handler.General
}

func InitHandlers(service *serviceInit) *HandlerInit {
	return &HandlerInit{
		Authentication: *handler.Account(&service.Account),
		General:        *handler.GeneralHandler(),
	}
}
