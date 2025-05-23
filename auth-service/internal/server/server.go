package server

import (
	"auth-service/internal/config"
	"auth-service/internal/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	Addr   string
}

func New(cfg config.Config) *Server {

	return &Server{
		router: gin.Default(),
		Addr:   fmt.Sprintf(":%s", cfg.ServerConfig.WebPort),
	}

}

func (s *Server) ListenAndServe() error {

	s.RegisterRoutes()
	s.RegisterMiddleWare()
	err := s.router.Run(s.Addr)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) RegisterRoutes() {
	s.router.GET("/", handlers.HealthCheck)
	s.router.POST("/auth/register", handlers.Registration)
}
func (s *Server) RegisterMiddleWare() {
	s.router.Use(handlers.LogMiddleWare())
}
