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

type LoginData struct {
	AccountID      string
	UserID         string
	CustomerID     string
	FirstName      string
	PrimaryAccount int
	AccountStatus  int
	TimeZone       string
	Password       string
}

type JWTPayload struct {
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
