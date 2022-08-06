package repository

import "github.com/gin-gonic/gin"
import "hb-backend-v1/model"

type AccountMock struct {
}

func (am AccountMock) Login(c *gin.Context, form model.LoginRequest) (bool, model.LoginDataResponse, string) {
	var result model.LoginDataResponse
	return false, result, ""
}

func (am AccountMock) Registration(c *gin.Context, form model.RegistrationRequest) (bool, string, string) {
	return false, "result", ""
}
