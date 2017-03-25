package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/techvein/gozen/models"
	"github.com/techvein/gozen/models/json"
)

// ユーザーコントローラー
type UserController struct{}

// 登録情報を返す
func (user UserController) Profile() gin.HandlerFunc {
	return jsonController(func(c *gin.Context) (interface{}, models.Error) {
		if models.LoginUser == nil {
			return nil, models.NewError(http.StatusBadRequest, "ログインしてください。")
		}
		return models.LoginUser.ToJson(), nil
	})
}

// registration idを登録する
func (user UserController) RegistrationID() gin.HandlerFunc {
	return jsonController(func(c *gin.Context) (interface{}, models.Error) {
		if models.LoginUser == nil {
			return nil, models.NewError(http.StatusBadRequest, "ログインしてください。")
		}

		regID := c.PostForm("registration_id")
		if regID == "" {
			return nil, models.NewError(http.StatusNotAcceptable, "not registraion_id key")
		}
		user := models.NewUser()
		if err := user.SaveRegistrationID(regID); err != nil {
			return nil, models.NewError(http.StatusInternalServerError, err.Error())
		}
		return json.MessageResponse{
			"saved registration id",
		}, nil
	})
}
