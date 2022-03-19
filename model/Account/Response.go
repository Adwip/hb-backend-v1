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

type LoginResult struct {
	Id       int
	Name     string
	Username string
	Email    string
	Password string
}

type AuthResult struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
	Token     string `json:"token"`
}
