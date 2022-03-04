package database

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
import "os"

var Connection *sql.DB

func InitDB(){
	// var err error
	connection, err := sql.Open("mysql",os.Getenv("MY_SQL_CONNECTION"))

	if err != nil{
		fmt.Println("Failed to connect DB", err)
	}

	if err = connection.Ping(); err != nil{
		fmt.Println("DB Unreachabel", err)
	}

	// model.DB = connection
	Connection = connection
}