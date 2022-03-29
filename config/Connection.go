package config

import "database/sql"

type ConnectionInt interface {
	InitConnection()
	GetConnection() *sql.DB
	TestConnection() (bool, error)
}
