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
/*
func Login(username string, email string)(LoginResult, error){
	md5 := md5.New()
	var result = LoginResult{}
	err := model.QueryRow("SELECT Id, Name, Username, Email FROM table WHERE Username = ? OR Email = ?", username, email)
	if err!=nil{
		return nil, err
	}
	return LoginResult, nil
}*/




