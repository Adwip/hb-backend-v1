package config

/*
import "database/sql"
import _ "github.com/lib/pq"
import "fmt"
import "os"

var postgreConnection *sql.DB

type PostgreSQL struct {
}

func PostgreDB() ConnectionIntf {
	database := &PostgreSQL{}
	return database
}

func (PostgreSQL) InitConnection() {
	// var err error
	connResult, err := sql.Open("mysql", os.Getenv("POSTGRE_URL"))
	// connResult, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/new_hubing_com")
	if err != nil {
		fmt.Println("Failed to connect Postgre DB", err)
	}

	if err = connResult.Ping(); err != nil {
		fmt.Println("Postgre DB Unreachabel", err)
	}

	// model.DB = connection
	// fmt.Println(connResult)
	postgreConnection = connResult
}

func (PostgreSQL) GetConnection() *sql.DB {
	return postgreConnection
}*/
