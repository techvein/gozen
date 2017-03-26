package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/techvein/gozen/models"
	"github.com/techvein/gozen/models/json"
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

// HTMLを返す共通コントローラー
func htmlController(fn func(*gin.Context) (interface{}, error), tmpl string) func(*gin.Context) {
	return func(c *gin.Context) {
		repository, err := fn(c)
		if err != nil {
			e, ok := err.(models.Error)
			if ok && e != nil {
				c.HTML(e.StatusCode(), "error.tmpl.html", e)
			}
			return
		}
		c.HTML(http.StatusOK, tmpl, repository)
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
