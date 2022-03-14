package database

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
import "os"

var connection *sql.DB

func InitDB(){
	// var err error
	connResult, err := sql.Open("mysql",os.Getenv("MY_SQL_CONNECTION"))

	if err != nil{
		fmt.Println("Failed to connect DB", err)
	}

	if err = connResult.Ping(); err != nil{
		fmt.Println("DB Unreachabel", err)
	}

	// model.DB = connection
	connection = connResult
}

func GetConnection() *sql.DB{
	return connection
}