package library

import accountForm "hb-backend-v1/model/account"

var header accountForm.JWTHeader
var payload accountForm.JWTPayload

type IdentityLib struct {
}

func Identity() *IdentityLib {
	identity := &IdentityLib{}
	return identity
}

func (IdentityLib) SetHeader(reqHeader *accountForm.JWTHeader) {
	header = *reqHeader
}

func (IdentityLib) SetPayload(reqPayload *accountForm.JWTPayload) {
	payload = *reqPayload
}

func (IdentityLib) GetUserID() string {
	return payload.UserID
}
