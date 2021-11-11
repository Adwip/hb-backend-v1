package database

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
import "hb-backend-v1/model"

var connection *sql.DB

func InitDB(){

	var err error
	connection, err = sql.Open("mysql","root:@tcp(127.0.0.1:3306)/new_hubing_com")

	if err != nil{
		fmt.Printf("Failed to connect DB", err)
	}

	if err = connection.Ping(); err != nil{
		fmt.Printf("DB Unreachabel", err)
	}

	model.DB = connection
}