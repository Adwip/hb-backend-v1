package model

import "database/sql"
import "hb-backend-v1/config/database"

func Query(query string, scan func(*sql.Rows) error) error{
	connect, err := database.Connect()

	if err != nil{
		return err
	}
	defer connect.Close()

	rows, err := connect.Query(query)

	if err != nil{
		return err
	}
	
	defer rows.Close()

	return scan(rows)
}

func Update() int {
	return 1
}

func Delete() int {
	return 1
}