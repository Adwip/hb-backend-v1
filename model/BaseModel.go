package model

import "database/sql"
import "fmt"
import "hb-backend-v1/config/database"
// import "reflect"

var DB *sql.DB

type Dao struct{
	Query string
	Rows *sql.Rows
}

func (dao *Dao) Select(param ...interface{}) error{
	var rows *sql.Rows
	var err error
	// DB = sqlConnection.GetConnection()
	DB = database.Connection
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
	DB = database.Connection
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

func (dao *Dao) Update(param ...interface{}) (bool, error) {
	if len(param)<1{
		return false, nil
	}
	
	// DB = sqlConnection.GetConnection()
	DB = database.Connection
	_, err := DB.Exec(dao.Query, param)

	if err != nil{
		return false, err
	}

	return true, nil
}

func (dao *Dao) Delete(param ...interface{}) (bool, error) {
	if len(param)<1{
		return false, nil
	}
	
	// DB = sqlConnection.GetConnection()
	DB = database.Connection
	_, err := DB.Exec(dao.Query, param)

	if err != nil{
		return false, err
	}

	return true, nil
}

func (dao *Dao) Insert(param ...interface{}) (bool, error) {
	if len(param)<1{
		return false, nil
	}

	// DB = sqlConnection.GetConnection()
	DB = database.Connection
	_, err := DB.Exec(dao.Query, param...)

	if err != nil{
		return false, err
	}

	return true, nil
}

