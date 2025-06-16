package v1

import (
	"api-gateway/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup, jwtSecret string) {
	g.GET("/", HealthCheck)
	gr := g.Group("/auth/")
	RegisterAuthRoutes(gr)
	userG := g.Group("/user/")
	RegisterUserRoutes(userG, jwtSecret)

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

func RegisterUserRoutes(g *gin.RouterGroup, jwtSecret string) {
	g.Use(middleware.AuthMiddleware(jwtSecret))
	g.GET("/users", FetchUsers)
}
