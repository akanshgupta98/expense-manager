package middleware

import (
	"github.com/akanshgupta98/go-logger"
	"github.com/gin-gonic/gin"
)

func LogMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Infof("API: %s, Method: %s, Host: %s", ctx.Request.URL, ctx.Request.Method, ctx.Request.Host)
		ctx.Next()
	}
}
