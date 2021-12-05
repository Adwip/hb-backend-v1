package auth

import _ "github.com/lestrrat-go/jwx/jwt"

type Auth struct{
	Uid			string
	Name		string
	Username	string
	Customer 	bool
	User 		bool
	Status		int
}

func (auth *Auth)GenerateToken(){

}

func (auth *Auth)ParseToken(){

}