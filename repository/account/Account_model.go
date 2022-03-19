package account

import "hb-backend-v1/repository"

// import _ "database/sql"
// import _ "crypto/md5"
import "hb-backend-v1/library/auth"
import "hb-backend-v1/library/authentication"

import "encoding/json"

// import _ "fmt"
import "hb-backend-v1/library/dateTime"

// import _ "github.com/google/uuid"
import accountForm "hb-backend-v1/model/account"
import "hb-backend-v1/model"

var Dao = repository.Dao{}

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

/*
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
}*/

func Login(unameMail string, password string) model.RepoResponse {
	// md5 := md5.New()
	var result accountForm.LoginResult
	timeNow := dateTime.DateTimeNow()

	Dao.Query = "select name, username, email, password from account where username = ? OR email = ?"
	exists, row, _ := Dao.SelectOne(unameMail, unameMail)
	if !exists {
		// return false, loginResult, error
		return model.RepoResponse{Success: false, Msg: "User not exists"}
	}

	row.Scan(&result.Name, &result.Username, &result.Email, &result.Password)

	if Oke := auth.VerifyPassword(password, result.Password); !Oke {
		// return false, loginResult, error
		return model.RepoResponse{Success: false, Msg: "User not exists"}
	}
	payloadJson := authentication.Payload{Id: result.Id, Name: result.Name, UserType: true, KeepLogin: true}
	payload, errJson := json.Marshal(payloadJson)
	if errJson != nil {
		// return false, loginResult, errJson
		return model.RepoResponse{Success: false, Msg: "Login rejected"}
	}
	token, errToken := authentication.GenerateToken("SHA256", "JWT", payload)
	if errToken != nil {
		// return false, loginResult, errToken
		return model.RepoResponse{Success: false, Msg: "Login rejected"}
	}

	// result := finalResult{result.Id, result.Name, true, timeNow, token}
	authResult := accountForm.AuthResult{
		Name:      result.Name,
		Username:  result.Username,
		CreatedAt: timeNow,
		Token:     token,
	}

	// return true, loginResult, nil
	return model.RepoResponse{Success: true, Data: authResult, Msg: "Login rejected"}
}

/*
func UpdatePassword(form UpdatePasswordForm)bool{
	query := "update account where "
	result, err := Dao.Update(username, email)
}*/

/*
func RegistrationUser(form accountForm.RegistrationForm) (bool, error) {
	// _ = form
	form.Password = authentication.SHA256encode(form.Password, "12345")
	id := uuid.New()
	Dao.Query = "INSERT INTO account (id, name, username, email, password) VALUES (?, ?, ?, ?, ?)"
	insert, err := Dao.QueryModifier(id, form.Name, form.Username, form.Email, form.Password)
	return insert, err
}*/
