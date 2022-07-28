package utils

type Authentication interface {
	GenerateJWT()
}

type authImp struct {
}

func AuthUtils() Authentication {
	return &authImp{}
}

func (authImp) GenerateJWT() {

}
