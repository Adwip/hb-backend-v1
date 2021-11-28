package model

import "database/sql"
import "fmt"
// import "reflect"

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

func (dao *Dao) SelectOne(param ...interface{}) (bool, *sql.Row, error){
	var result, resultCheck *sql.Row
	var exists bool
	var err error
	queryCheck := fmt.Sprintf("SELECT exists (%s)", dao.Query)
	if len(param) > 0{
		resultCheck = DB.QueryRow(queryCheck, param...)
		result = DB.QueryRow(dao.Query, param...)
	}else{
		resultCheck = DB.QueryRow(queryCheck)
		result = DB.QueryRow(dao.Query)
	}
	err = resultCheck.Scan(&exists)
	if !exists {
		return false, result, err
	}
	/*
	if err != nil && err != sql.ErrNoRows{
		return false, result, err
	}*/
	
	return true, result, nil
}

func Update(query string) error {
	return nil
}

func Delete() int {
	return 1
}

func Insert() {
	
}

