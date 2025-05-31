package server

import (
	"auth-service/internal/config"
	"auth-service/internal/handlers"
	"fmt"

	"github.com/akanshgupta98/go-logger"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	Addr   string
}

func New(cfg config.Config) *Server {
	logger.Debugf("Server new called")
	gin.SetMode(gin.ReleaseMode)

	return &Server{
		router: gin.New(),
		Addr:   fmt.Sprintf(":%s", cfg.ServerConfig.WebPort),
	}

}

func (s *Server) ListenAndServe() error {

	s.RegisterMiddleWare()
	s.RegisterRoutes()
	err := s.router.Run(s.Addr)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) RegisterRoutes() {
	s.router.GET("/", handlers.HealthCheck)
	s.router.POST("/auth/users", handlers.Registration)
	s.router.GET("/auth/users", handlers.FetchAllUsers)
	s.router.POST("/auth/login", handlers.Login)
}
func (s *Server) RegisterMiddleWare() {
	// s.router.Use(handlers.CORSMiddleWare())
	s.router.Use(handlers.LogMiddleWare())
}
