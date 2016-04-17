package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gozen/models"
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