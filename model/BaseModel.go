package model

import "database/sql"
// import "fmt"
// import "reflect"

var DB *sql.DB

type Dao struct{
	Query string
	Rows *sql.Rows
	Row *sql.Row
}

func Query(query string, scan func(*sql.Rows) error) error{
	/*
	connect, err := database.Connect()

	if err != nil{
		return err
	}
	defer connect.Close()
	*/
	
	rows, err := DB.Query(query)

	if err != nil{
		return err
	}
	
	defer rows.Close()

	return scan(rows)
}

func (dao *Dao) Select(param ...interface{}) error{
	var rows *sql.Rows
	var err error
	if len(param) > 0{
		rows, err = DB.Query(dao.Query, param...)
	}else{
		rows, err = DB.Query(dao.Query)
	}

	if err != nil{
		defer rows.Close()
		return err
	}
	dao.Rows = rows
	
	return nil
}

func (dao *Dao) SelectOne(param ...interface{}) error{
	var row *sql.Row
	
	if len(param) > 0{
		row = DB.QueryRow(dao.Query, param...)
	}else{
		row = DB.QueryRow(dao.Query)
	}
	
	dao.Row = row
	return nil
}

func Update(query string) error {
	return nil
}

func Delete() int {
	return 1
}

func Insert() {
	
}