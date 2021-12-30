package account

type AccountScan struct{
	Id 				int
	Name 			string
	Username 		string
	Email			string
}

type LoginResult struct{
	Id				int
	Name			string
	Username		string
	Email			string
	Password		string
}


type LoginForm struct{
	Username string `json:"username"`
	Email	string	`json:"email"`
	Password string	`json:"password"`
	KeepLogin bool	`json:"keepLogin"`
}

type jwtPayload struct{
	Id			int
	Name		string
	UserType	bool
}

type finalResult struct{
	Id			int
	Name		string
	UserType	bool
	LoginTime	int
	ValidUntil	int
	Token		string


}