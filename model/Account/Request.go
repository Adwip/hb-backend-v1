package account

type LoginForm struct {
	UnameMail string `json:"unameMail"`
	Password  string `json:"password"`
	KeepLogin bool   `json:"keepLogin"`
}

type RegistrationForm struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Username    string `json:"username"`
	TimeZone    string `json:"timeZone"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	AccountType int8   `json:"accountType"`
}

type UpdatePasswordForm struct {
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}
