package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup) {
	g.GET("/", HealthCheck)
	gr := g.Group("/auth/")
	RegisterAuthRoutes(gr)

}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Health": "OK",
	})
}

func RegisterAuthRoutes(gr *gin.RouterGroup) {
	gr.POST("/login", Login)
}
