package sqlConnect

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import  "fmt"

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

func connect() (*sql.DB, error){
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/hubing_db")

	if err != nil{
		return nil, err
	}
	return db, nil
}

func Query() ([]resultDB, error){
	connect, err := connect()
	fmt.Println("Koneksi berhasil")

	if err != nil{
		fmt.Println("Koneksi gagal")
		return nil, err
	}

	defer connect.Close()

	rows, err := connect.Query("SELECT * FROM alamat_store")

	defer rows.Close()

	if err != nil{
		fmt.Println("Query gagal")
		return nil, err
	}

	var result []resultDB

	for rows.Next(){
		var each = resultDB{}
		var err = rows.Scan(&each.Id_alamat_store, &each.Id_tk, &each.Alamat, &each.Id_wil, &each.Latitude, &each.Longitude)
		if err != nil{
			fmt.Println("Mapping gagal", err.Error())
			return result, nil
		}
		result = append(result, each)
	}
	fmt.Println(result)
	return result, nil
}