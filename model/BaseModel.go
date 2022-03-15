package model

import (
	"database/sql"
	"hb-backend-v1/config/database"
)

// import "reflect"

var DB *sql.DB

type Dao struct {
	Query string
	Rows  *sql.Rows
}

func InitDao(query string) Dao {
	return Dao{
		Query: query,
	}
}

func (dao *Dao) Select(param ...interface{}) error {
	var rows *sql.Rows
	var err error
	DB = database.GetConnection() /*
		if len(param) > 0{
			rows, err = DB.Query(dao.Query, param...)
		}else{
			rows, err = DB.Query(dao.Query)
		}*/
	rows, err = DB.Query(dao.Query, param...)
	defer rows.Close()

	if err != nil {
		return err
	}

	dao.Rows = rows

	return nil
}

func (dao *Dao) SelectOne(param ...interface{}) (*sql.Rows, error) {
	// var result, resultCheck *sql.Row
	// var result *sql.Row
	// var exists bool
	// var err error
	DB = database.GetConnection()
	// DB = database.Connection
	// queryCheck := fmt.Sprintf("SELECT exists (%s)", dao.Query)
	// if len(param) > 0{
	// resultCheck = DB.QueryRow(queryCheck, param...)
	// result = DB.QueryRow(dao.Query, param...)
	rows, err := DB.Query(dao.Query, param...)
	// }else{
	// resultCheck = DB.QueryRow(queryCheck)
	// result = DB.QueryRow(dao.Query)
	// }
	// err = resultCheck.Scan(&exists)
	// if !exists {
	// return false, result, err
	// }
	if err != nil {
		return nil, err
	}
	return rows, nil
	/*
		if err != nil && err != sql.ErrNoRows{
			return false, result, err
		}*/

	// return true, result, nil
}

func (dao *Dao) QueryModifier(param ...interface{}) (bool, error) {
	if len(param) < 1 {
		return false, nil
	}
	DB = database.GetConnection()
	_, err := DB.Exec(dao.Query, param...)

	if err != nil {
		return false, err
	}

	return true, nil
}
