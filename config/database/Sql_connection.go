package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var connection *sql.DB

// type DB struct {
// 	connection *sql.DB
// }

// func (db *DB) Init(){
// 	//koneksi
// }

// func (db *DB) GetConnection() *sql.DB{
// 	return db.connection
// }

// func GetConfiguration() *DB{
// 	return &DB
// }

func InitDB() {
	// var err error
	connResult, err := sql.Open("mysql", os.Getenv("MY_SQL_CONNECTION"))

	if err != nil {
		fmt.Println("Failed to connect DB", err)
	}

	if err = connResult.Ping(); err != nil {
		fmt.Println("DB Unreachabel", err)
	}

	// model.DB = connection
	connection = connResult
}

func GetConnection() *sql.DB {
	return connection
}
