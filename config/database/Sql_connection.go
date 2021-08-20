package sqlConnect

import "database/sql"

type conectionString struct{
	dbname string
	username string
	password string
	host string
}

func connect() (*sql.DB, error){
	db, err := sql.Open("mysql","root:@tcp(localhost3306)/hubing_db")

	if err!=nil{
		return nil, err
	}
	return db, nil
}

func Query() (*sql.Rows, error){
	connect, err := connect()

	if err != nil{
		return nil, err
	}

	defer connect.Close()

	rows, err := connect.Query("select * from alamat_store")

	if err != nil{
		return nil,err
	}

	return rows, nil
}