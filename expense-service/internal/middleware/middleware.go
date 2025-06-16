package middleware

import (
	"github.com/akanshgupta98/go-logger/v2"
	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		logger.Infof("API: %s Method: %s Host: %s ", c.Request.URL, c.Request.Method, c.Request.Host)
		c.Next()
	}
}
