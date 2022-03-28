package repository

import "encoding/json"
import "fmt"
import "hb-backend-v1/library"
import _ "hb-backend-v1/library/authentication"
import _ "hb-backend-v1/library/dateTime"
import "hb-backend-v1/model"
import accountForm "hb-backend-v1/model/account"
import "database/sql"
import "hb-backend-v1/config"
import "github.com/gin-gonic/gin"
import "context"
import "time"
import "os"
import _ "reflect"
import "github.com/google/uuid"

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
	conn *sql.DB
}

func Account() *AccountObj {
	database := config.Database()
	connSring := database.GetConnection()
	account := &AccountObj{
		conn: connSring,
	}
	return account
}

func (account *AccountObj) Login(c *gin.Context, form *accountForm.LoginForm) *model.RepoResponse {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	jwtLib := library.JsonWT()
	hash := library.Hash()
	time := library.Time()
	passwordKey := os.Getenv("PASSWORD_SECRET_KEY")
	jwtKey := os.Getenv("JWT_SECRET_KEY")
	var result accountForm.LoginResult

	defer cancel()
	sqlStatement := "select id AS userID, firstName, 1 AS primaryAccount, 1 AS accountStatus, timeZone, password from account inner join account_information on account.id = account_information.id_account where username = ? OR email = ?"
	query := account.conn.QueryRowContext(ctx, sqlStatement, form.UnameMail, form.UnameMail)
	err := query.Scan(&result.UserID, &result.FirstName, &result.PrimaryAccount, &result.AccountStatus, &result.TimeZone, &result.Password)
	if err == sql.ErrNoRows {
		return &model.RepoResponse{Success: false, Msg: "User not exists | no result"}
	} else if err != nil {
		return &model.RepoResponse{Success: false, Msg: err.Error()}
	}
	currentDateTime := time.CurrentTimeUnix()
	// utc := time.CurrentTimeUTC()
	// dbFormat := time.CurrentDateTimeDbFormat()
	// fmt.Println(currentDateTime)
	// fmt.Println(utc)
	// fmt.Println(dbFormat)
	// fmt.Println(os.Getenv("PASSWORD_KEY"))
	// fmt.Println(passwordKey)
	// fmt.Println(reflect.TypeOf(passwordKey))
	if Approved := hash.VerifyPassword(form.Password, result.Password, passwordKey); !Approved {
		return &model.RepoResponse{Success: false, Msg: "User not exists | password is wrong"}
	}

	JWTPayload := accountForm.JWTPayload{
		UserID:         result.UserID,
		FirstName:      result.FirstName,
		PrimaryAccount: result.PrimaryAccount,
		AccountStatus:  result.AccountStatus,
		TimeZone:       result.TimeZone,
		CreatedAt:      currentDateTime,
	}

	payload, errJson := json.Marshal(JWTPayload)
	if errJson != nil {
		return &model.RepoResponse{Success: false, Msg: "Failed to generate token"}
	}
	token, errToken := jwtLib.GenerateToken("SHA256", "JWT", payload, jwtKey)
	if errToken != nil {
		return &model.RepoResponse{Success: false, Msg: "Login rejected"}
	}

	authResponse := accountForm.AuthResponse{
		UserID:         result.UserID,
		FirstName:      result.FirstName,
		PrimaryAccount: result.PrimaryAccount,
		AccountStatus:  result.AccountStatus,
		TimeZone:       result.TimeZone,
		CreatedAt:      currentDateTime,
		Token:          token,
	}

	return &model.RepoResponse{Success: true, Data: authResponse}
}

/*
func UpdatePassword(form UpdatePasswordForm)bool{
	query := "update account where "
	result, err := Dao.Update(username, email)
}*/

func (account AccountObj) RegistrationUser(c *gin.Context, form accountForm.RegistrationForm) *model.RepoResponse {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	// _ = form
	var accountTypeTable string
	var execErr3 error
	// var errorMsg map[string]string
	id := uuid.New()
	accInfID := uuid.New()
	hash := library.Hash()
	timeLib := library.Time()
	passwordKey := os.Getenv("PASSWORD_SECRET_KEY")
	createdAt := timeLib.CurrentDateTimeDbFormat()
	// _ = id
	// _ = ctx
	defer cancel()
	// sqlStatement := "insert into "
	accountTable := "insert into account (id, username, email, primaryAccount, password) values(?, ?, ?, ?, ?)"
	accountInformationTable := "insert into account_information (id_AccInf, id_account, firstName, lastName, timeZone, phone, createdAt) values(?, ?, ?, ?, ?, ?, ?)"
	if form.AccountType == 1 {
		accountTypeTable = "insert into user (id_account, registeredAt, status) values(?, ?, ?)"
	} else {
		accountTypeTable = "insert into user (id_account, registeredAt, status) values(?, ?, ?)"
	}
	hashedPassword := hash.SHA256(form.Password, passwordKey)
	// _ = hashedPassword

	tx, err := account.conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return &model.RepoResponse{Success: false, Msg: err.Error()}
	}

	_, execErr1 := tx.Exec(accountTable, id, form.Username, form.Email, form.AccountType, hashedPassword)
	if execErr1 != nil {
		tx.Rollback()
		// fmt.Println(execErr1)
		// return &model.RepoResponse{Success: true, Msg: execErr1.Error()}
	}
	_, execErr2 := tx.Exec(accountInformationTable, accInfID, id, form.FirstName, form.LastName, form.TimeZone, form.Phone, createdAt)
	if execErr2 != nil {
		// fmt.Println(execErr2)
		tx.Rollback()
		// return &model.RepoResponse{Success: true, Msg: execErr2.Error()}
	}

	if form.AccountType == 1 {
		_, execErr3 = tx.Exec(accountTypeTable, id, createdAt, 1)
	} else {
		_, execErr3 = tx.Exec(accountTypeTable, id, createdAt, 1)
	}
	if execErr3 != nil {
		// fmt.Println(execErr3)
		tx.Rollback()
		// return &model.RepoResponse{Success: true, Msg: execErr3.Error()}
	}
	errTrans := tx.Commit()

	if errTrans != nil {
		return &model.RepoResponse{Success: false, Msg: "Failed"}
	}

	return &model.RepoResponse{Success: true}
}
