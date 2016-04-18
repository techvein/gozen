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
		return commonOauthController(c, githubUser)
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
		return commonOauthController(c, googleUser)
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
