package service

import "github.com/gin-gonic/gin"
import "hb-backend-v1/model"
import "testing"
import "hb-backend-v1/repository"
import "net/http/httptest"
import "fmt"

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
