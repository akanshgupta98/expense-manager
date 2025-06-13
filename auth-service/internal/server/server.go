package server

import (
	"auth-service/internal/config"
	v1 "auth-service/internal/handlers/v1"
	"auth-service/internal/middleware"
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
	r := gin.New()
	r.Use(middleware.LogMiddleWare())

	return &Server{
		router: r,
		Addr:   fmt.Sprintf(":%s", cfg.ServerConfig.WebPort),
	}

}

func (s *Server) ListenAndServe() error {

	r := s.router.Group("/auth/v1")
	v1.RegisterRoutes(r)
	err := s.router.Run(s.Addr)
	if err != nil {
		return err
	}

	return nil
}
