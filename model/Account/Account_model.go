package account

import "hb-backend-v1/model"
import _"database/sql"
import _"crypto/md5"
import "hb-backend-v1/library/auth"
import tokenGen "hb-backend-v1/library/authentication"
import "encoding/json"
import _"fmt"
import "hb-backend-v1/library/dateTime"

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


func Login(username string, email string, password string)(bool, finalResult, error){
	// md5 := md5.New()
	var result LoginResult
	var finalRest finalResult
	timeNow := dateTime.DateTimeNow()
	Dao.Query = "select id, name, username, email, password from account where username = ? OR email = ?"
	exists, row, error := Dao.SelectOne(username, email)
	if !exists {
		return false, finalRest, error
	}

	row.Scan(&result.Id, &result.Name, &result.Username, &result.Email, &result.Password)

	if Oke := auth.VerifyPassword(password, result.Password); !Oke{
		return false, finalRest, error
	}
	payloadJson := tokenGen.Payload{Id:result.Id, Name:result.Name, UserType:true, KeepLogin:true}
	payload, errJson := json.Marshal(payloadJson)
	if errJson!=nil{
		return false, finalRest, errJson
	}
	token, errToken := tokenGen.GenerateToken("SHA256", "JWT", payload)
	if errToken != nil{
		return false, finalRest, errToken
	}

	// result := finalResult{result.Id, result.Name, true, timeNow, token}
	finalRest.Id = result.Id
	finalRest.Name = result.Name
	finalRest.UserType = true
	finalRest.CreatedAt = timeNow
	finalRest.Token = token


	return true, finalRest, nil
}




