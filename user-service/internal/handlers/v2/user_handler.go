package v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup) {

	g.GET("/", HealthCheck)
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health":     "Ok",
		"APIVersion": "v2",
	})
}
