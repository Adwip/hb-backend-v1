package database

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

type conectionString struct{
	dbname string
	username string
	password string
	host string
}

type resultDB struct{
	Id_alamat_store int
	Id_tk 			int
	Alamat 			string
	Id_wil 			int
	Latitude 		sql.NullString
	Longitude 		sql.NullString
}

func Connect() (*sql.DB, error){
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/hubing_db")
	if err != nil{
		return nil, err
	}
	return db, nil
}

func Query() ([]resultDB, error){
	connect, err := Connect()

	if err != nil{
		return nil, err
	}

	defer connect.Close()

	rows, err := connect.Query("SELECT * FROM alamat_store")

	defer rows.Close()

	if err != nil{
		return nil, err
	}

	fmt.Println("Hasil unScan query Database")
	fmt.Println(rows)
	var result []resultDB

	for rows.Next(){
		var each = resultDB{}
		var err = rows.Scan(&each.Id_alamat_store, &each.Id_tk, &each.Alamat, &each.Id_wil, &each.Latitude, &each.Longitude)
		if err != nil{
			return result, nil
		}
		result = append(result, each)
	}
	return result, nil
}