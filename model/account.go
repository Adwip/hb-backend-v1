package model

type LoginRequest struct {
	UnameMail string `json:"unameMail"`
	Password  string `json:"password"`
	KeepLogin bool   `json:"keepLogin"`
}

type RegistrationRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Username    string `json:"username"`
	TimeZone    string `json:"timeZone"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	AccountType int8   `json:"accountType"`
}

type UpdatePasswordRequest struct {
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type AccountResponse struct {
	Id       int
	Name     string
	Username string
	Email    string
}

type LoginResponse struct {
	Id       int
	Name     string
	Username string
	Email    string
	Password string
}

type FinalResultResponse struct {
	Id        int
	Name      string
	UserType  bool
	CreatedAt string
	Token     string
}

/*
type AuthResult struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	TimeZone  string `json:"timeZone"`
	CreatedAt int64  `json:"createdAt"`
	Token     string `json:"token"`
}*/

type LoginDataResponse struct {
	AccountID      string
	UserID         string
	CustomerID     string
	FirstName      string
	PrimaryAccount int
	AccountStatus  int
	TimeZone       string
	Password       string
}

type JWTPayloadResponse struct {
	AccountID      string `json:"accountID"`
	UserExists     bool   `json:"userAccount"`
	UserID         string `json:"userID"`
	CustomerExists bool   `json:"customerAccount"`
	CustomerID     string `json:"customerID"`
	FirstName      string `json:"firstName"`
	PrimaryAccount int    `json:"primaryAccount"`
	AccountStatus  int    `json:"accountStatus"`
	TimeZone       string `json:"timeZone"`
	CreatedAt      int64  `json:"createdAt"`
}

type AuthResponse struct {
	AccountID      string `json:"accountID"`
	UserID         string `json:"userID"`
	CustomerID     string `json:"customerID"`
	FirstName      string `json:"firstName"`
	PrimaryAccount int    `json:"primaryAccount"`
	AccountStatus  int    `json:"accountStatus"`
	TimeZone       string `json:"timeZone"`
	CreatedAt      int64  `json:"createdAt"`
	Token          string `json:"token"`
}

type JWTHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
