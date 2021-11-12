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


func Login(username string, email string)(LoginResult, error){
	// md5 := md5.New()
	var result LoginResult
	Dao.Query = "select id, name, username, email from account where username = ? OR email = ?"
	err := Dao.SelectOne(username, email)
	
	if err != nil{
		if Dao.Row.Scan(&result.Id).ErrNoRows{

		}
		return result, err
	}
	
	Dao.Row.Scan(&result.Id, &result.Name, &result.Username, &result.Email)


	return result, err
}




