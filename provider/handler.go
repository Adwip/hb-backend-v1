package provider

import handler "hb-backend-v1/controller"

type handlerInit struct {
	Authentication handler.AccountCtrl
}

func InitHandlers(service *serviceInit) *handlerInit {
	return &handlerInit{
		Authentication: *handler.Account(&service.Authentication),
	}
}
