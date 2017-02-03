package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gozen/config"
	"gozen/models"
	"gozen/oauth"
	"log"
)

// 認証コントローラー
type OAuthController struct {
	oauth.User
}

func NewOauthController(user oauth.User) OAuthController {
	return OAuthController{user}
}

// ログインへリダイレクトを行う
func (self OAuthController) Login() gin.HandlerFunc {
	return jsonController(func(c *gin.Context) (interface{}, models.Error) {
		if self.User.GetClientID() == nil {
			return nil, models.NewError(http.StatusInternalServerError, "client_id has not been set.")
		}
		if self.User.GetClientSecret() == nil {
			return nil, models.NewError(http.StatusInternalServerError, "client_secret has not been set.")
		}
		url := self.User.GenerateLoginUrl()
		c.Redirect(http.StatusTemporaryRedirect, url)
		return nil, nil
	})
}

// CallBack処理を行う
func (self OAuthController) CallBack() gin.HandlerFunc {
	return jsonController(func(c *gin.Context) (interface{}, models.Error) {
		log.Println(c.Request.URL)
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

	log.Println(url)
	c.Redirect(http.StatusTemporaryRedirect, url)
	return nil, nil
}
