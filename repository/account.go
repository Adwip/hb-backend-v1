package repository

// import "encoding/json"
import "fmt"
import "hb-backend-v1/utils"
import "hb-backend-v1/model"
import "database/sql"
import "github.com/gin-gonic/gin"
import "context"
import "time"
import "os"
import "github.com/google/uuid"

// import "hb-backend-v1/config"

type AccountObj struct {
	conn *sql.DB
}

type AccountInt interface {
	Login(*gin.Context, model.LoginRequest) (bool, model.LoginDataResponse, string)
	Registration(*gin.Context, model.RegistrationRequest) (bool, string, string)
	// UpdatePassword(*gin.Context, model.UpdatePasswordRequest) *model.RepoResponse
}

func Account(db *sql.DB) AccountInt {
	return &AccountObj{
		conn: db,
	}
}

func (account *AccountObj) Login(c *gin.Context, form model.LoginRequest) (bool, model.LoginDataResponse, string) {
	var result model.LoginDataResponse

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	hash := utils.Hash()
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

func (account AccountObj) Registration(c *gin.Context, form model.RegistrationRequest) (bool, string, string) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	var accountTypeTable string

	id := uuid.New()
	accInfID := uuid.New()
	generalID := uuid.New()
	timeLib := utils.Time()
	createdAt := timeLib.CurrentDateTimeDbFormat()
	defer cancel()

	accountTable := "insert into account (id_account, username, email, primaryAccount, password) values(?, ?, ?, ?, ?)"
	accountInformationTable := "insert into account_information (id_account_inf, account, firstName, lastName, timeZone, phone, createdAt) values(?, ?, ?, ?, ?, ?, ?)"
	if form.AccountType == 1 {
		accountTypeTable = "insert into user (id_user, account, registeredAt, status) values(?, ?, ?, ?)"
	} else {
		accountTypeTable = "insert into customer (id_customer, account, registeredAt, status) values(?, ?, ?, ?)"
	}

	tx, err := account.conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		fmt.Println(err)
		return false, "", "Failed to add"
	}

	_, execErr1 := tx.Exec(accountTable, id, form.Username, form.Email, form.AccountType, form.Password)
	if execErr1 != nil {
		tx.Rollback()
		fmt.Println(execErr1)
	}
	_, execErr2 := tx.Exec(accountInformationTable, accInfID, id, form.FirstName, form.LastName, form.TimeZone, form.Phone, createdAt)
	if execErr2 != nil {
		fmt.Println(execErr2)
		tx.Rollback()
	}

	_, execErr3 := tx.Exec(accountTypeTable, generalID, id, createdAt, 1)

	if execErr3 != nil {
		fmt.Println(execErr3)
		tx.Rollback()
	}
	errTrans := tx.Commit()

	if errTrans != nil {
		fmt.Println(errTrans)
		return false, "", "Failed to add"
	}
	return true, id.String(), "Failed to add"
}

/*
func (account AccountObj) UpdatePassword(c *gin.Context, form model.UpdatePasswordRequest) *model.RepoResponse {
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
}*/
