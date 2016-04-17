package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gozen/models"
	"gozen/models/json"
)

func middleWares(engine *gin.Engine) {
	engine.RouterGroup.Use(
		TokenMiddleWare(),
	)
}

// JSONを返す共通コントローラー
func jsonController(fn func(*gin.Context) (interface{}, models.Error)) func(*gin.Context) {
	return func(c *gin.Context) {

		withCORS(c)

		repository, err := fn(c)
		if err != nil {
			errorJSON := json.ErrorJSON{
				Code:    err.StatusCode(),
				Message: err.Message(),
			}
			c.JSON(errorJSON.Code, errorJSON)
			return
		}
		c.JSON(http.StatusOK, repository)
	}
}

func noImpl(c *gin.Context) (interface{}, models.Error) {
	return nil, models.NewError(http.StatusNotFound, "no implemt")

}

// CORS ヘッダーを出力する
func withCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Keep-Alive,User-Agent,If-Modified-Since,Cache-Control,Content-Type,Authorization")
	c.Header("Access-Control-Max-Age", "1728000")
}