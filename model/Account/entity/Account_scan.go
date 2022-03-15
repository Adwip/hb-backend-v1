package entity

type AccountScan struct {
	Id       int
	Name     string
	Username string
	Email    string
}

type LoginResult struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type LoginForm struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	KeepLogin bool   `json:"keepLogin"`
}

type RegistrationForm struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type finalResult struct {
	Id        int
	Name      string
	UserType  bool
	CreatedAt string
	Token     string
}

type UpdatePasswordForm struct {
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}
