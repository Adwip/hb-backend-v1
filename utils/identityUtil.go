package utils

import _ "context"
import m "hb-backend-v1/model"
import "github.com/gin-gonic/gin"

// import "fmt"

type IdentityLib struct {
	header        m.JWTHeaderResponse
	headerExists  bool
	payload       m.JWTPayloadResponse
	payloadExists bool
}

func Identity(ctx *gin.Context) *IdentityLib {
	header, headerExists := ctx.Get("JWTHeader")
	payload, payloadExists := ctx.Get("JWTPayload")
	identity := &IdentityLib{
		header:        header.(m.JWTHeaderResponse),
		headerExists:  headerExists,
		payload:       payload.(m.JWTPayloadResponse),
		payloadExists: payloadExists,
	}
	return identity
}

func (id *IdentityLib) GetUserID() string {
	if id.payloadExists {
		return id.payload.UserID
	}
	return ""
}

func (id *IdentityLib) GetAccountID() string {
	if id.payloadExists {
		return id.payload.AccountID
	}
	return ""
}

func (id *IdentityLib) GetCustomerID() string {
	if id.payloadExists {
		return id.payload.CustomerID
	}
	return ""
}

func (id *IdentityLib) GetFirstname() string {
	if id.payloadExists {
		return id.payload.FirstName
	}
	return ""
}

func (id *IdentityLib) GetTimezone() string {
	if id.payloadExists {
		return id.payload.TimeZone
	}
	return ""
}

func (id *IdentityLib) IsAuthenticated() bool {
	return id.payloadExists
}
