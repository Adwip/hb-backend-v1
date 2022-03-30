package config

import "database/sql"

type ConnectionIntf interface {
	InitConnection()
	GetConnection() *sql.DB
}
