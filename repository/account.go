package repository

// import "encoding/json"
import "fmt"
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

func (account *AccountObj) Login(c *gin.Context, form *accountForm.LoginForm) (bool, accountForm.LoginData, string) {
	var result accountForm.LoginData

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	hash := library.Hash()
	passwordKey := os.Getenv("PASSWORD_SECRET_KEY")

	defer cancel()
	sqlStatement := "SELECT id_account AS accountID, IF(id_user!='', id_user, '') as userID, IF(id_customer!='', id_customer, '') as customerID, firstName, 1 AS primaryAccount, 1 AS accountStatus, timeZone, password FROM account INNER JOIN account_information ON account.id_account = account_information.account LEFT JOIN user ON account.id_account=user.account LEFT JOIN customer ON account.id_account=customer.account where username = ? OR email = ?"
	query := account.conn.QueryRowContext(ctx, sqlStatement, form.UnameMail, form.UnameMail)

	err := query.Scan(&result.AccountID, &result.UserID, &result.CustomerID, &result.FirstName, &result.PrimaryAccount, &result.AccountStatus, &result.TimeZone, &result.Password)
	if err == sql.ErrNoRows {
		return false, result, "Account not exists | no result"
	} else if err != nil {
		return false, result, "Account not exists | no result"
	}

	if Approved := hash.VerifyPassword(form.Password, result.Password, passwordKey); !Approved {
		return false, result, "User not exists | password is wrong"
	}
	return true, result, ""
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
	accountTable := "insert into account (id_account, username, email, primaryAccount, password) values(?, ?, ?, ?, ?)"
	accountInformationTable := "insert into account_information (id_account_inf, account, firstName, lastName, timeZone, phone, createdAt) values(?, ?, ?, ?, ?, ?, ?)"
	if form.AccountType == 1 {
		accountTypeTable = "insert into user (account, registeredAt, status) values(?, ?, ?)"
	} else {
		accountTypeTable = "insert into customer (account, registeredAt, status) values(?, ?, ?)"
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
		fmt.Println(execErr1)
		// return &model.RepoResponse{Success: true, Msg: execErr1.Error()}
	}
	_, execErr2 := tx.Exec(accountInformationTable, accInfID, id, form.FirstName, form.LastName, form.TimeZone, form.Phone, createdAt)
	if execErr2 != nil {
		fmt.Println(execErr2)
		tx.Rollback()
		// return &model.RepoResponse{Success: true, Msg: execErr2.Error()}
	}

	_, execErr3 := tx.Exec(accountTypeTable, id, createdAt, 1)

	if execErr3 != nil {
		fmt.Println(execErr3)
		tx.Rollback()
		// return &model.RepoResponse{Success: true, Msg: execErr3.Error()}
	}
	errTrans := tx.Commit()

	if errTrans != nil {
		return &model.RepoResponse{Success: false, Msg: errTrans.Error()}
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
