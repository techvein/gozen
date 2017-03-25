package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/techvein/gozen/oauth"
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
			userGroup.POST("/registrationid", userController.RegistrationID())
		}

		pushGroup := api.Group("push")
		{
			pushNotificationController := PushNotificationController{}
			pushGroup.POST("/send", pushNotificationController.Send())
		}

	}

	// OAuth関連
	oauthGroup := r.Group("oauth")
	{
		github := NewOauthController(oauth.NewOAuthGitHub())
		oauthGroup.GET("/github-login", github.Login())
		oauthGroup.GET("/github_cb", github.CallBack())

		google := NewOauthController(oauth.NewOAuthGoogle())
		oauthGroup.GET("/google-login", google.Login())
		oauthGroup.GET("/google_cb", google.CallBack())

		facebook := NewOauthController(oauth.NewOAuthFacebook())
		oauthGroup.GET("/facebook-login", facebook.Login())
		oauthGroup.GET("/facebook_cb", facebook.CallBack())
	}

	return r

}
