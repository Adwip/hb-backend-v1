package account

type Account struct {
	Id       int
	Name     string
	Username string
	Email    string
}

type Login struct {
	Id       int
	Name     string
	Username string
	Email    string
	Password string
}

type FinalResult struct {
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

type LoginResult struct {
	UserID         string
	FirstName      string
	PrimaryAccount int
	AccountStatus  int
	TimeZone       string
	Password       string
}

type JWTPayload struct {
	UserID         string `json:"userID"`
	FirstName      string `json:"firstName"`
	PrimaryAccount int    `json:"primaryAccount"`
	AccountStatus  int    `json:"accountStatus"`
	TimeZone       string `json:"timeZone"`
	CreatedAt      int64  `json:"createdAt"`
}

type AuthResponse struct {
	UserID         string `json:"userID"`
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
