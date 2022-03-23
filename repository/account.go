package repository

import "encoding/json"
import "fmt"
import "hb-backend-v1/library"
import "hb-backend-v1/library/authentication"
import "hb-backend-v1/library/dateTime"

import "hb-backend-v1/model"
import accountForm "hb-backend-v1/model/account"

import "database/sql"
import "hb-backend-v1/config"

// import _ "database/sql"
// import _ "crypto/md5"
// import _ "github.com/google/uuid"
// import _ "fmt"

// var Dao = repository.Dao{}

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

type AccountObj struct {
	// baseRepo *DaoRepo
	conn *sql.DB
}

func Account() *AccountObj {
	// repoHandler := Dao()
	database := config.Database()
	connSring := database.GetConnection()
	account := &AccountObj{
		conn: connSring,
	}
	return account
}

func (account *AccountObj) Login(unameMail string, password string) model.RepoResponse {
	var result accountForm.LoginResult
	var jwtLib = library.JWT{}
	timeNow := dateTime.DateTimeNow()

	// account.baseRepo.Query = "select id AS userID, firstName, 1 AS primaryAccount, 1 AS accountStatus, password from account inner join account_information on account.id = account_information.id_account where username = ? OR email = ?"
	// query := account.baseRepo.SelectOne(unameMail, unameMail)
	sqlStatement := "select id AS userID, firstName, 1 AS primaryAccount, 1 AS accountStatus, password from account inner join account_information on account.id = account_information.id_account where username = ? OR email = ?"
	query := account.conn.QueryRow(sqlStatement, unameMail, unameMail)
	err := query.Scan(&result.UserID, &result.FirstName, &result.PrimaryAccount, &result.AccountStatus, &result.Password)
	fmt.Println(err)
	if err == sql.ErrNoRows || err == sql.ErrConnDone {
		return model.RepoResponse{Success: false, Msg: "User not exists | no result"}
	}
	// fmt.Println(sql.ErrConnDone)
	// fmt.Println(sql.ErrNoRows)
	// fmt.Println(account.conn)
	// fmt.Println(result)
	// fmt.Println("Database => ", result.Password)
	// fmt.Println("User => ", password)

	// row.Scan(&result.UserID, &result.FirstName, &result.PrimaryAccount, &result.AccountStatus, &result.Password)

	if Approved := authentication.PasswordVerification(password, result.Password); !Approved {
		return model.RepoResponse{Success: false, Msg: "User not exists | password is wrong"}
	}

	JWTPayload := accountForm.JWTPayload{
		UserID:         result.UserID,
		FirstName:      result.FirstName,
		PrimaryAccount: result.PrimaryAccount,
		AccountStatus:  result.AccountStatus,
		CreatedAt:      timeNow,
	}

	payload, errJson := json.Marshal(JWTPayload)
	if errJson != nil {
		return model.RepoResponse{Success: false, Msg: "Failed to generate token"}
	}
	// token, errToken := authentication.GenerateToken("SHA256", "JWT", payload)
	token, errToken := jwtLib.GenerateToken("SHA256", "JWT", payload)
	if errToken != nil {
		return model.RepoResponse{Success: false, Msg: "Login rejected"}
	}

	authResponse := accountForm.AuthResponse{
		UserID:         result.UserID,
		FirstName:      result.FirstName,
		PrimaryAccount: result.PrimaryAccount,
		AccountStatus:  result.AccountStatus,
		CreatedAt:      timeNow,
		Token:          token,
	}

	// return true, loginResult, nil
	return model.RepoResponse{Success: true, Data: authResponse, Msg: "Login rejected"}
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
