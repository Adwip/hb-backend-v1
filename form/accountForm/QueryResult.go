package accountForm

type Account struct{
	Id 				int
	Name 			string
	Username 		string
	Email			string
}

type Login struct{
	Id				int
	Name			string
	Username		string
	Email			string
	Password		string
}

type FinalResult struct{
	Id			int
	Name		string
	UserType	bool
	CreatedAt	string
	Token		string
}

type LoginResult struct{
	Id				int
	Name			string
	Username		string
	Email			string
	Password		string
}