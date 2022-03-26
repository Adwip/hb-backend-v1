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

func (id *IdentityLib) GetFirstname() string {
	if id.payloadExists {
		return id.payload.FirstName
	}
	return ""
}