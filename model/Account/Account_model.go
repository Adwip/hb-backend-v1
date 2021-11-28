package account

import "hb-backend-v1/model"
import _"database/sql"
import _"crypto/md5"

var Dao = model.Dao{}
/*
func AllAccount() ([]AccountScan, error){
	var result []AccountScan
	err := model.Query("select id, name, username, email from account",func(rows *sql.Rows)error{
		for rows.Next(){
			var each = AccountScan{}
			var err = rows.Scan(&each.Id, &each.Name, &each.Username, &each.Email)
			if err != nil{
				return err
			}
			result = append(result, each)
		}
		return nil
	})
	if err != nil{
		return nil,err
	}
	return result,nil
}*/

func AllAccount() ([]AccountScan, error){
	var result []AccountScan
	Dao.Query = "select id, name, username, email from account"
	err := Dao.Select()
	if err != nil{
		return nil,err
	}
	for Dao.Rows.Next(){
		var each = AccountScan{}
		var err = Dao.Rows.Scan(&each.Id, &each.Name, &each.Username, &each.Email)
		if err != nil{
			return nil, err
		}
		result = append(result, each)
	}
	defer Dao.Rows.Close()

	return result,nil
}


func Login(username string, email string)(bool, LoginResult, error){
	// md5 := md5.New()
	var result LoginResult
	Dao.Query = "select id, name, username, email from account where username = ? OR email = ?"
	exists, row, error := Dao.SelectOne(username, email)
	if !exists {
		return false, result, error
	}

	row.Scan(&result.Id, &result.Name, &result.Username, &result.Email)

	return true, result, nil
}




