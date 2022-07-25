package config

import "database/sql"

type MySQLIntf interface {
	InitMySQL() *sql.DB
}
