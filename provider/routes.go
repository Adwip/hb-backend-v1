package provider

import "github.com/gin-gonic/gin"

func InitRoutes(handler *handlerInit, router *gin.Engine) {
	handler.Authentication.Routes(router)
} /*
type status int

const (
	CreateStatus status = iota + 1
	FinishStatus
	CancelStatus
)


func (s status) Value() {

}

func updateStatus(toStatus status) error {

}

func sd () {
	var ads int = 2
	CreateStatus.Value()
	updateStatus(ads)
}*/
