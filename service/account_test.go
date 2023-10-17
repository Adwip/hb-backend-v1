package service

import (
	"fmt"
	"hb-backend-v1/model"
	"hb-backend-v1/repository"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var authRepoMock = repository.AccountMock{}

func Test_Login(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	authService := AccountService{authRepoMock}
	var formMock model.LoginRequest

	success, _, _ := authService.Login(c, formMock)
	if !success {
		t.Fail()
	}
	fmt.Println("Pengetesan")
}
