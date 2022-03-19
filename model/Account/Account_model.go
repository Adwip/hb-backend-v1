package account

import "hb-backend-v1/model"
import _ "database/sql"
import _ "crypto/md5"
import "hb-backend-v1/library/auth"
import "hb-backend-v1/library/authentication"
import "encoding/json"
import _ "fmt"
import "hb-backend-v1/library/dateTime"
import "hb-backend-v1/form/accountForm"
import "github.com/google/uuid"

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

func AllAccount() ([]AccountScan, error) {
	var result []AccountScan
	Dao.Query = "select id, name, username, email from account"
	err := Dao.Select()
	if err != nil {
		return nil, err
	}
	for Dao.Rows.Next() {
		var each = AccountScan{}
		var err = Dao.Rows.Scan(&each.Id, &each.Name, &each.Username, &each.Email)
		if err != nil {
			return nil, err
		}
		result = append(result, each)
	}
	defer Dao.Rows.Close()

	return result, nil
}

func Login(unameMail string, password string) (bool, map[string]interface{}, error) {
	// md5 := md5.New()
	var result accountForm.LoginResult
	var loginResult = make(map[string]interface{})
	timeNow := dateTime.DateTimeNow()
	Dao.Query = "select id, name, username, email, password from account where username = ? OR email = ?"
	exists, row, error := Dao.SelectOne(unameMail, unameMail)
	if !exists {
		return false, loginResult, error
	}

	row.Scan(&result.Id, &result.Name, &result.Username, &result.Email, &result.Password)

	if Oke := auth.VerifyPassword(password, result.Password); !Oke {
		return false, loginResult, error
	}
	payloadJson := authentication.Payload{Id: result.Id, Name: result.Name, UserType: true, KeepLogin: true}
	payload, errJson := json.Marshal(payloadJson)
	if errJson != nil {
		return false, loginResult, errJson
	}
	token, errToken := authentication.GenerateToken("SHA256", "JWT", payload)
	if errToken != nil {
		return false, loginResult, errToken
	}

	// result := finalResult{result.Id, result.Name, true, timeNow, token}
	loginResult["id"] = result.Id
	loginResult["name"] = result.Name
	loginResult["userType"] = true
	loginResult["createdAt"] = timeNow
	loginResult["token"] = token

	return true, loginResult, nil
}

/*
func UpdatePassword(form UpdatePasswordForm)bool{
	query := "update account where "
	result, err := Dao.Update(username, email)
}*/

func RegistrationUser(form accountForm.RegistrationForm) (bool, error) {
	// _ = form
	form.Password = authentication.SHA256encode(form.Password, "12345")
	id := uuid.New()
	Dao.Query = "INSERT INTO account (id, name, username, email, password) VALUES (?, ?, ?, ?, ?)"
	insert, err := Dao.QueryModifier(id, form.Name, form.Username, form.Email, form.Password)
	return insert, err
}
