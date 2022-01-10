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

type RegistrationForm struct {
	Username	string	`json:"username"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
	UserType	uint8	`json:"userType"`
}

type finalResult struct{
	Id			int
	Name		string
	UserType	bool
	CreatedAt	string
	Token		string
}