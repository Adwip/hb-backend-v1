package account

type LoginForm struct{
	Username string `json:"username"`
	Email	string	`json:"email"`
	Password string	`json:"password"`
	KeepLogin bool	`json:"keepLogin"`
}