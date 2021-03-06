package config

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
import "os"

var connection *sql.DB

type MySQL struct {
}

func Database() ConnectionIntf {
	database := &MySQL{}
	return database
}

func (MySQL) InitConnection() {
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

func (MySQL) GetConnection() *sql.DB {
	return connection
}
