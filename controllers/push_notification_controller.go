package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-gcm"

	"gozen/config"
	"gozen/models"
	"gozen/models/json"
)

type PushNotificationController struct{}

// push送信
func (push PushNotificationController) Send() gin.HandlerFunc {
	return jsonController(func(c *gin.Context) (interface{}, models.Error) {

		// TODO: Perhaps, you should change the key.
		body := c.PostForm("body")
		if body == "" {
			return nil, models.NewError(http.StatusNotAcceptable, "body key should be posted。")
		}

		if !models.LoginUser.RegID.Valid {
			return nil, models.NewError(http.StatusNotFound, "registration id isn't registered")
		}

		m := gcm.HttpMessage{
			To:           models.LoginUser.RegID.String,
			Notification: &gcm.Notification{Body: body},
		}
		_, sendErr := gcm.SendHttp(config.Push.Gcm.ApiKey, m) // responseは使わない
		if sendErr != nil {
			return nil, models.NewError(http.StatusInternalServerError, sendErr.Error())
		}

		return json.MessageResponse{
			"sent push notificaiton",
		}, nil

	})
}
