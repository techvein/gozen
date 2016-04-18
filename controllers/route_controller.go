package controllers

import (
	"github.com/gin-gonic/gin"
)

// 参考: https://github.com/konjoot/microservice_experiments/blob/master/gin-gonic/src/router/router.go
func Routes() *gin.Engine {
	r := gin.Default()

	middleWares(r)

	// TODO: とりあえずルーティングはここに書いていくことにしましょうか。
	// map[string]gin.HandlerFuncに詰め込んでループみたいな感じにしてもよさそう
	api := r.Group("api")
	{
		userGroup := api.Group("user")
		{
			userController := UserController{}
			userGroup.GET("/profile", userController.Profile())
		}

	}

	// OAuth関連
	oauthGroup := r.Group("oauth")
	{
		oauthController := OAuthController{}
		oauthGroup.GET("/github-login", oauthController.GithubLogin())
		oauthGroup.GET("/github_cb", oauthController.GithubCallBack())

		oauthGroup.GET("/google-login", oauthController.GoogleLogin())
		oauthGroup.GET("/google_cb", oauthController.GoogleCallBack())
	}

	return r

}
