package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup) {
	g.GET("/", HealthCheck)
	gr := g.Group("/auth/")
	RegisterAuthRoutes(gr)
	userG := g.Group("/user/")
	RegisterUserRoutes(userG)

}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Health": "OK",
	})
}

func RegisterAuthRoutes(gr *gin.RouterGroup) {
	gr.POST("/login", Login)
	gr.POST("/register", Register)
}

func RegisterUserRoutes(g *gin.RouterGroup) {
	g.GET("/users", FetchUsers)
}
