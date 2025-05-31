package server

import (
	"api-gateway/internal/config"
	v1 "api-gateway/internal/handlers/v1"
	"api-gateway/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func New(cfg config.Config) *Server {
	server := Server{
		mux:  gin.New(),
		addr: fmt.Sprintf(":%s", cfg.ServerCfg.WebPort),
	}
	return &server
}

func (s *Server) ListenAndServe() error {
	rg := s.mux.Group("/api/v1")
	s.mux.Use(middleware.CORSMiddleWare())
	middleware.RegisterMiddleware(rg)
	v1.RegisterRoutes(rg)

	err := s.mux.Run(s.addr)
	if err != nil {
		return err
	}
	return nil
}
