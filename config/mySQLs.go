package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	_ "os"
)

// var connection *sql.DB

type MySQL struct {
	url string
}

func Database(url string) MySQLIntf {
	return &MySQL{
		url: url,
	}
}

func (ms MySQL) InitMySQL() *sql.DB {
	// var err error
	connResult, err := sql.Open("mysql" /*os.Getenv("MY_SQL_URL")*/, ms.url)
	// connResult, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/new_hubing_com")
	if err != nil {
		fmt.Println("Failed to connect DB", err)
	}

	if err = connResult.Ping(); err != nil {
		fmt.Println("DB Unreachabel", err)
	}

	// model.DB = connection
	// fmt.Println(connResult)
	// connection = connResult
	return connResult
}
