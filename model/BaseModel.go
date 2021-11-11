package model

import "database/sql"
import _"fmt"

var DB *sql.DB

type Dao struct{
	Query string
	Rows *sql.Rows
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

func (dao *Dao) Select() error{

	result, err := DB.Query(dao.Query)

	if err != nil{
		defer result.Close()
		return err
	}
	dao.Rows = result
	
	return nil
}

func Update(query string) error {
	return nil
}

func Delete() int {
	return 1
}