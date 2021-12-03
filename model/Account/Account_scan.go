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
	Token			string
}