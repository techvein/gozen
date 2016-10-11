package controllers_test

import (
	"net/http"
	"testing"
	"time"

	"gozen/controllers"
	"gozen/db"
	"gozen/models"
	. "gozen/models/json"
)

// ユーザーの確認
func TestUser(t *testing.T) {
	var userJson UserJson
	var mr MessageResponse

	const sessionToken = "1234"

	// テスト用にユーザーを追加
	testUser := &testUser{}
	user, lastInsertAuthId, err := testUser.registerNewUser(sessionToken)
	if err != nil {
		t.Error(err)
	}

	// 未ログイン時のプロフィール確認
	userProfileNotLoginTest := commonData{
		filePath:    "user_profile_not_login.json",
		requestPath: "profile",
		exp:         mr,
		act:         mr,
		statusCode:  http.StatusBadRequest,
	}
	checkGozen(t, "user", userProfileNotLoginTest)

	// ログイン済みのときのプロフィール確認
	userProfileTest := commonData{
		filePath:     "user_profile.json",
		requestPath:  "profile",
		exp:          userJson,
		act:          userJson,
		statusCode:   http.StatusOK,
		sessionToken: sessionToken,
	}
	checkGozen(t, "user", userProfileTest)

	// Authを元に戻す
	deleteBuilder := db.GetSession().
		DeleteFrom(models.NewAuth().TableName()).
		Where("id = ?", lastInsertAuthId)
	deleteBuilder.Exec()

	// Userを元に戻す
	deleteBuilder = db.GetSession().
		DeleteFrom(models.NewUser().TableName()).
		Where("id = ?", user.Id.Int64)
	deleteBuilder.Exec()

}

type testUser struct{}

func (testUser *testUser) registerNewUser(token string) (user *models.User, lastInsertAuthId int64, err error) {
	auth := &models.Auth{
		Source:   "github",
		SourceId: 6789,
		Email:    "abc@example.com",
	}

	user = &models.User{
		Name:         "def",
		Email:        auth.Email,
		SessionToken: token,
	}

	user.TokenExpire.Time = time.Now().Add(time.Hour * 24 * 30)
	user.LastLoginAt.Time = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	result, err := db.GetSession().InsertInto(models.NewUser().TableName()).Columns(
		"id",
		"name",
		"email",
		"session_token",
		"token_expire",
		"last_login_at",
	).Record(
		user,
	).Exec()

	if err != nil {
		return nil, 0, err
	}

	lastInsertUserId, _ := result.LastInsertId()
	user.Id.Int64 = lastInsertUserId
	auth.UserId = uint64(lastInsertUserId)

	result, err = db.GetSession().InsertInto(auth.TableName()).Columns(
		"id",
		"user_id",
		"source",
		"source_id",
		"email",
	).Record(
		auth,
	).Exec()

	if err != nil {
		return nil, 0, err
	}

	lastInsertAuthId, _ = result.LastInsertId()
	return user, lastInsertAuthId, nil
}

func BenchmarkUserProfile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := controllers.UserController{}
		c.Profile()
	}
}
