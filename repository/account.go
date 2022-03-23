package repository

import "encoding/json"
import _ "fmt"
import "hb-backend-v1/library"
import "hb-backend-v1/library/authentication"
import "hb-backend-v1/library/dateTime"
import "hb-backend-v1/model"
import accountForm "hb-backend-v1/model/account"
import "database/sql"
import "hb-backend-v1/config"
import "github.com/gin-gonic/gin"
import "context"
import "time"

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
	conn   *sql.DB
	ctx    *context.Context
	cancel context.CancelFunc
}

func Account(c *gin.Context) *AccountObj {
	database := config.Database()
	connSring := database.GetConnection()
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	account := &AccountObj{
		conn:   connSring,
		ctx:    &ctx,
		cancel: cancel,
	}
	return account
}

func (account *AccountObj) Login(unameMail string, password string) *model.RepoResponse {
	var result accountForm.LoginResult
	var jwtLib = library.JWT{}
	timeNow := dateTime.DateTimeNow()
	defer account.cancel() //=> Error is from here, but I need to call cancel() which as field in the Account struct
	// fmt.Println(*account.cancel)
	sqlStatement := "select id AS userID, firstName, 1 AS primaryAccount, 1 AS accountStatus, password from account inner join account_information on account.id = account_information.id_account where username = ? OR email = ?"
	query := account.conn.QueryRowContext(*account.ctx, sqlStatement, unameMail, unameMail)
	err := query.Scan(&result.UserID, &result.FirstName, &result.PrimaryAccount, &result.AccountStatus, &result.Password)
	if err == sql.ErrNoRows {
		return &model.RepoResponse{Success: false, Msg: "User not exists | no result"}
	} else if err != nil {
		return &model.RepoResponse{Success: false, Msg: err.Error()}
	}
	// fmt.Println(err)
	if Approved := authentication.PasswordVerification(password, result.Password); !Approved {
		return &model.RepoResponse{Success: false, Msg: "User not exists | password is wrong"}
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
		return &model.RepoResponse{Success: false, Msg: "Failed to generate token"}
	}
	token, errToken := jwtLib.GenerateToken("SHA256", "JWT", payload)
	if errToken != nil {
		return &model.RepoResponse{Success: false, Msg: "Login rejected"}
	}

	authResponse := accountForm.AuthResponse{
		UserID:         result.UserID,
		FirstName:      result.FirstName,
		PrimaryAccount: result.PrimaryAccount,
		AccountStatus:  result.AccountStatus,
		CreatedAt:      timeNow,
		Token:          token,
	}

	return &model.RepoResponse{Success: true, Data: authResponse}
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
