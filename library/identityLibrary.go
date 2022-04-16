package library

import _ "context"
import accountForm "hb-backend-v1/model/account"
import "github.com/gin-gonic/gin"

type IdentityLib struct {
	header        accountForm.JWTHeader
	headerExists  bool
	payload       accountForm.JWTPayload
	payloadExists bool
}

func Identity(ctx *gin.Context) *IdentityLib {
	header, headerExists := ctx.Get("JWTHeader")
	payload, payloadExists := ctx.Get("JWTPayload")
	identity := &IdentityLib{
		header:        header.(accountForm.JWTHeader),
		headerExists:  headerExists,
		payload:       payload.(accountForm.JWTPayload),
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
