package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/techvein/gozen/models"
)

const (
	SESSION_TOKEN_KEY = "X-Session-Token"
)

func TokenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(SESSION_TOKEN_KEY)
		models.NewUser().FindUserBySessionToken(token)

		c.Next()
	}
}
