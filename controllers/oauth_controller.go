package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/logger"

	"gozen/config"
	"gozen/models"
	"gozen/oauth"
)

// 認証コントローラー
type oAuthController struct {
	oauth.User
}

func NewOauthController(user oauth.User) oAuthController{
	return oAuthController{user}
}

// ログインへリダイレクトを行う
func (self oAuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := self.User.GenerateLoginUrl()
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

// CallBack処理を行う
func (self oAuthController) CallBack() gin.HandlerFunc {
	return jsonController(func(c *gin.Context) (interface{}, models.Error) {
		logger.Infoln(c.Request.URL)
		user, err := self.User.Callback(c.Query("state"), c.Query("code"))
		if err != nil {
			return nil, models.NewError(http.StatusBadRequest, err.Error())
		}
		return commonOauthController(c, user)
	})
}

func commonOauthController(c *gin.Context, user oauth.User) (interface{}, models.Error) {
	token, err := models.NewUser().Auth(user)
	if err != nil {
		return nil, models.NewError(http.StatusBadRequest, err.Error())
	}
	url := fmt.Sprintf("%s?token=%s", config.Oauth.AfterOauthUrl, token)

	logger.Infoln(url)
	c.Redirect(http.StatusTemporaryRedirect, url)
	return nil, nil
}
