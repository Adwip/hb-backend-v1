package authentication

type Header struct{
	Alg string
	Typ string
}

type Payload struct{
	Id			int
	Name		string
	UserType	bool
	KeepLogin	bool
}