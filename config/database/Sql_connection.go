package database

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

var Connection *sql.DB

func InitDB(){
	// var err error
	connection, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/new_hubing_com")

	if err != nil{
		fmt.Println("Failed to connect DB", err)
	}

	if err = connection.Ping(); err != nil{
		fmt.Println("DB Unreachabel", err)
	}

	// model.DB = connection
	Connection = connection
}