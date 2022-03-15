package account

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"hb-backend-v1/form/accountForm"
	"hb-backend-v1/library/auth"
	"hb-backend-v1/library/authentication"
	"hb-backend-v1/library/dateTime"
	"hb-backend-v1/model"
	"hb-backend-v1/model/account/entity"

	"github.com/google/uuid"
)

type AccountModel struct {
	DB *sql.DB
}

func InitAccountModel(database *sql.DB) AccountModel {
	return AccountModel{
		DB: database,
	}
}

func (am *AccountModel) AllAccount() ([]entity.AccountScan, error) {
	var result []entity.AccountScan
	var dao = model.InitDao("select id, name, username, email from account")
	// Dao.Query =
	err := dao.Select()
	if err != nil {
		return nil, err
	}
	for dao.Rows.Next() {
		var each = entity.AccountScan{}
		var err = dao.Rows.Scan(&each.Id, &each.Name, &each.Username, &each.Email)
		if err != nil {
			return nil, err
		}
		result = append(result, each)
	}
	defer dao.Rows.Close()

	return result, nil
}

func (am *AccountModel) Login(unameMail string, password string) (bool, map[string]interface{}, error) {
	// md5 := md5.New()
	var result []entity.LoginResult
	var loginResult = make(map[string]interface{})
	timeNow := dateTime.DateTimeNow()
	// Dao.Query =
	query := "select id, name, username, email, password from account where username = $1 OR email = $1"
	// var parameter []interface{}
	// if unameMail != "" {
	// 	parameter = append(parameter, unameMail)
	// }
	rows, err := am.DB.Query(query, unameMail)
	if err != nil {
		return false, loginResult, err
	}
	// row, error := dao.SelectOne(parameter)
	// if !exists {
	// return false, loginResult, error
	// }

	fmt.Println("error")

	err = rows.Scan(&result)
	if err != nil {
		return false, loginResult, err
	}

	if oke := auth.VerifyPassword(password, result.Password); !oke {
		return false, loginResult, nil
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

// func UpdatePassword(form UpdatePasswordForm)bool{
// 	query := "update account where "
// 	result, err := Dao.Update(username, email)
// }

func (am *AccountModel) RegistrationUser(form accountForm.RegistrationForm) (bool, error) {
	// _ = form
	form.Password = authentication.SHA256encode(form.Password, "12345")
	id := uuid.New()
	dao := model.InitDao("INSERT INTO account (id, name, username, email, password) VALUES (?, ?, ?, ?, ?)")
	// Dao.Query =
	insert, err := dao.QueryModifier(id, form.Name, form.Username, form.Email, form.Password)
	return insert, err
}
