package accountForm

type LoginForm struct{
	UnameMail string `json:"unameMail"`
	Password string	`json:"password"`
	KeepLogin bool	`json:"keepLogin"`
}

type RegistrationForm struct {
	Name		string	`json:"name"`
	Username	string	`json:"username"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}

type UpdatePasswordForm struct{
	OldPassword 	string	`json:"oldPassword"`
	NewPassword 	string	`json:"newPassword"`
	ConfirmPassword string	`json:"confirmPassword"`
}