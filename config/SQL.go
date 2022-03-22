package config

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
import "os"

var connection *sql.DB

type DatabaseCfg struct {
}

func Database() *DatabaseCfg {
	database := &DatabaseCfg{}
	return database
}

func (DatabaseCfg) InitDB() {
	// var err error
	connResult, err := sql.Open("mysql", os.Getenv("MY_SQL_URL"))
	// connResult, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/new_hubing_com")
	if err != nil {
		fmt.Println("Failed to connect DB", err)
	}

	if err = connResult.Ping(); err != nil {
		fmt.Println("DB Unreachabel", err)
	}

	// model.DB = connection
	// fmt.Println(connResult)
	connection = connResult
}

func (DatabaseCfg) GetConnection() *sql.DB {
	return connection
}

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
