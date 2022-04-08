package repository

import "encoding/json"
import _ "fmt"
import "hb-backend-v1/library"
import "hb-backend-v1/model"
import accountForm "hb-backend-v1/model/account"
import "database/sql"
import "hb-backend-v1/config"
import "github.com/gin-gonic/gin"
import "context"
import "time"
import "os"
import "github.com/google/uuid"

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
	sqlStatement := "select id_account AS userID, firstName, 1 AS primaryAccount, 1 AS accountStatus, timeZone, password from account inner join account_information on account.id_account = account_information.account where username = ? OR email = ?"
	query := account.conn.QueryRowContext(ctx, sqlStatement, form.UnameMail, form.UnameMail)
	err := query.Scan(&result.UserID, &result.FirstName, &result.PrimaryAccount, &result.AccountStatus, &result.TimeZone, &result.Password)
	if err == sql.ErrNoRows {
		return &model.RepoResponse{Success: false, Msg: "User not exists | no result"}
	} else if err != nil {
		return &model.RepoResponse{Success: false, Msg: err.Error()}
	}
	currentDateTime := time.CurrentTimeUnix()

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

func (account AccountObj) RegistrationUser(c *gin.Context, form accountForm.RegistrationForm) *model.RepoResponse {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	// _ = form
	var accountTypeTable string
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
		accountTypeTable = "insert into customer (id_account, registeredAt, status) values(?, ?, ?)"
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

	_, execErr3 := tx.Exec(accountTypeTable, id, createdAt, 1)

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

func (account AccountObj) UpdatePassword(c *gin.Context, form accountForm.UpdatePasswordForm) *model.RepoResponse {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	identity := library.Identity(c)
	hash := library.Hash()
	passwordKey := os.Getenv("PASSWORD_SECRET_KEY")
	defer cancel()
	_ = ctx

	statement := "update account set password = ? where id = ? and password = ?"
	hashedPass := hash.SHA256(form.NewPassword, passwordKey)
	confirmHashedPass := hash.SHA256(form.ConfirmPassword, passwordKey)
	oldPassword := hash.SHA256(form.OldPassword, passwordKey)
	if hashedPass != confirmHashedPass {
		return &model.RepoResponse{Success: false, Msg: "Password confirm not matched"}
	}
	result, err := account.conn.ExecContext(ctx, statement, hashedPass, identity.GetUserID(), oldPassword)
	if err != nil {
		return &model.RepoResponse{Success: false, Msg: err.Error()}
	}
	rows, errAff := result.RowsAffected()
	if errAff != nil {
		return &model.RepoResponse{Success: false, Msg: errAff.Error()}
	}

	if rows < 1 {
		return &model.RepoResponse{Success: false, Msg: "Failed to update password"}
	}
	return &model.RepoResponse{Success: true}
}
