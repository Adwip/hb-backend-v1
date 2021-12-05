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