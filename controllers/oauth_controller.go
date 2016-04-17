package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/logger"

	"gozen/models"
	"gozen/oauth"
)

// TODO 暫定的に定数化 環境別設定にしたい
const after_auth_url = "http://example.com/auth?token="

// 認証コントローラー
type OAuthController struct{}

// Githubログインへリダイレクトを行う
func (self OAuthController) GithubLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := oauth.NewOAuthGitHub().GenerateLoginUrl()
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

// Github からのCallBack処理を行う
func (self OAuthController) GithubCallBack() gin.HandlerFunc {
	return jsonController(func(c *gin.Context) (interface{}, models.Error) {
		githubUser, err := oauth.NewOAuthGitHub().Callback(c.Query("state"), c.Query("code"))

		if err != nil {
			return nil, models.NewError(http.StatusBadRequest, err.Error())
		}

		token, err := models.NewUser().Auth(githubUser)

		url := after_auth_url + token

		logger.Infoln(url)
		c.Redirect(http.StatusTemporaryRedirect, url)

		return nil, nil
	})
}

// Googleログインへリダイレクトを行う
func (self OAuthController) GoogleLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := oauth.NewOAuthGoogle().GenerateLoginUrl()
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

// Google からのCallBack処理を行う
func (self OAuthController) GoogleCallBack() gin.HandlerFunc {
	return jsonController(func(c *gin.Context) (interface{}, models.Error) {
		googleUser, err := oauth.NewOAuthGoogle().Callback(c.Query("state"), c.Query("code"))

		if err != nil {
			return nil, models.NewError(http.StatusBadRequest, err.Error())
		}

		token, err := models.NewUser().Auth(googleUser)

		url := after_auth_url + token

		logger.Infoln(url)
		c.Redirect(http.StatusTemporaryRedirect, url)

		return nil, nil
	})
}