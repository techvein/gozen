package models

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gocraft/dbr"

	"github.com/techvein/gozen/db"
	"github.com/techvein/gozen/models/json"
	"github.com/techvein/gozen/oauth"
)

type User struct {
	Id           dbr.NullInt64  `db:"id"`
	Name         string         `db:"name"`
	Email        string         `db:"email"`
	SessionToken string         `db:"session_token"`
	TokenExpire  dbr.NullTime   `db:"token_expire"`
	RegID        dbr.NullString `db:"registration_id"`
	LastLoginAt  dbr.NullTime   `db:"last_login_at"`
	CreatedAt    dbr.NullTime
	UpdatedAt    dbr.NullTime
}

var userColumns []string

var LoginUser *User

func init() {
	userColumns = NewUser().createColumns()
}

func (self *User) TableName() string {
	return "users"
}

func NewUser() *User {
	return new(User)
}

// 認証する
func (self *User) Auth(ou oauth.User) (string, error) {
	token := self.sessionId()

	auth := NewAuth()
	err := db.GetSession().Select("*").From(auth.TableName()).
		Where("source_id = ?", ou.GetID()).
		LoadStruct(auth)

	if err != nil {
		// 見つからない場合は新規登録
		if err == dbr.ErrNotFound {
			user, err := self.registerNewUser(ou, token)
			if err != nil {
				return "", err
			}
			return user.SessionToken, nil
		}
		return "", err
	}

	// SESSIONを更新する
	_, err = db.GetSession().Update(self.TableName()).Set(
		"session_token",
		token,
	).Where(
		"id = ?",
		auth.UserId,
	).Exec()

	err = db.GetSession().Select("*").From(self.TableName()).
		Where("id = ?", auth.UserId).
		LoadStruct(&self)

	return self.SessionToken, nil
}

// 新規ユーザーを登録する
func (self *User) registerNewUser(ou oauth.User, token string) (*User, error) {
	tx, _ := db.GetSession().Begin()
	defer tx.RollbackUnlessCommitted()

	var name, email string
	// TODO: name,email以外のプロパティを追加したくなったら、このif文を増やさないといけなくなるので、なんとかしたい
	if ou.GetName() != nil {
		name = *ou.GetName()
	}
	if ou.GetEmail() != nil {
		email = *ou.GetEmail()
	}

	auth := &Auth{
		Source:    ou.GetSource(),
		SourceId:  *ou.GetID(),
		Email:     email,
		CreatedAt: dbr.NewNullTime(time.Now()),
		UpdatedAt: dbr.NewNullTime(time.Now()),
	}

	user := &User{
		Name:         name,
		Email:        auth.Email,
		SessionToken: token,
		TokenExpire:  dbr.NewNullTime(time.Now().Add(time.Hour * 24 * 30)),
		LastLoginAt:  dbr.NewNullTime(time.Now()),
		CreatedAt:    dbr.NewNullTime(time.Now()),
		UpdatedAt:    dbr.NewNullTime(time.Now()),
	}

	result, err := db.GetSession().InsertInto(self.TableName()).Columns(
		"id",
		"name",
		"email",
		"session_token",
		"token_expire",
		"last_login_at",
		"created_at",
		"updated_at",
	).Record(
		user,
	).Exec()

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	lastInsertId, _ := result.LastInsertId()
	user.Id = dbr.NewNullInt64(lastInsertId)
	auth.UserId = uint64(lastInsertId)

	_, err = db.GetSession().InsertInto(auth.TableName()).Columns(
		"id",
		"user_id",
		"source",
		"source_id",
		"email",
		"created_at",
		"updated_at",
	).Record(
		auth,
	).Exec()

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	return user, nil
}

// https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/06.2.md
func (self *User) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// Jsonに変換する
func (self *User) ToJson() *json.UserJson {
	userJson := json.NewUserJson()

	userJson.Id = self.Id.Int64
	userJson.Name = self.Name
	userJson.Email = self.Email
	userJson.SessionToken = self.SessionToken

	return userJson
}

// セッショントークンをキーにUserを検索する
func (self *User) FindUserBySessionToken(token string) error {
	log.Println(token)
	err := db.GetSession().Select(userColumns...).From(self.TableName()).
		Where("session_token = ?", token).
		LoadStruct(self)
	if err != nil {
		LoginUser = nil
		log.Println(err)
		return err
	}
	log.Println(self.Id)
	LoginUser = self
	return nil
}

func (self *User) SaveRegistrationID(regID string) error {
	if LoginUser == nil {
		return NewError(http.StatusNotAcceptable, "you must login")
	}
	_, err := db.GetSession().Update(self.TableName()).Set(
		"registration_id", regID,
	).Where(
		dbr.Eq(self.TableName()+".id", LoginUser.Id.Int64),
	).Exec()
	if err != nil {
		return err
	}
	return nil
}

func (self *User) createColumns() []string {
	columns := make([]string, 0)

	columnSlices := [][]string{
		StructTagToColumns(NewUser()),
	}

	for _, columnSlices := range columnSlices {
		columns = append(columns, columnSlices...)
	}
	return columns

}
