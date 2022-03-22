package repository

import "database/sql"

// import "fmt"
import "hb-backend-v1/config"

// import "reflect"

// var DB *sql.DB

// database := config.Database()
type DaoRepo struct {
	Query      string
	Rows       *sql.Rows
	connection *sql.DB
}

func Dao() *DaoRepo {
	database := config.Database()
	dao := &DaoRepo{
		connection: database.GetConnection(),
	}
	return dao
}

func (dao *DaoRepo) Select(param ...interface{}) error {
	var rows *sql.Rows
	var err error
	db := dao.connection
	if len(param) > 0 {
		rows, err = db.Query(dao.Query, param...)
	} else {
		rows, err = db.Query(dao.Query)
	}

	if err != nil {
		defer rows.Close()
		return err
	}
	dao.Rows = rows

	return nil
}

func (dao *DaoRepo) SelectOne(param ...interface{}) *sql.Row {
	var result *sql.Row
	// var exists bool
	// var err error
	db := dao.connection
	// DB = config.Connection
	// queryCheck := fmt.Sprintf("SELECT exists (%s)", dao.Query)
	if len(param) > 0 {
		// resultCheck = db.QueryRow(queryCheck, param...)
		result = db.QueryRow(dao.Query, param...)
	} else {
		// resultCheck = db.QueryRow(queryCheck)
		result = db.QueryRow(dao.Query)
	}
	// err = resultCheck.Scan(&exists)
	// fmt.Println(exists)
	// if result == sql.ErrNoRows {
	// return false, result, err
	// }
	/*
		if err != nil && err != sql.ErrNoRows{
			return false, result, err
		}*/

	return result
}

func (dao *DaoRepo) QueryModifier(param ...interface{}) (bool, error) {
	if len(param) < 1 {
		return false, nil
	}
	db := dao.connection
	_, err := db.Exec(dao.Query, param...)

	if err != nil {
		return false, err
	}

	return true, nil
}
