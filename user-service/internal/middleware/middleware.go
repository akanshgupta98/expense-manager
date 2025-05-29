package middleware

import (
	"net/http"

	"github.com/akanshgupta98/go-logger/v2"
	"github.com/gin-gonic/gin"
)

func Log() gin.HandlerFunc {

	return func(c *gin.Context) {

		logger.Debugf("API Request: %s Method: %s Host: %s", c.Request.URL, c.Request.Method, c.Request.Host)
		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
			return
		}
		c.Next()
	}
}
