package provider

import "hb-backend-v1/repository"
import "database/sql"

type repositoryInit struct {
	Account repository.Account
}

func InitRepositories(db *sql.DB) *repositoryInit {
	return &repositoryInit{
		Account: repository.NewAccountRepo(db),
	}
}
