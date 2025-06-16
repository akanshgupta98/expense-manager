package server

import (
	"expense-service/internal/config"
	v1 "expense-service/internal/handlers/v1"
	"expense-service/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func New(cfg config.Config) *Server {
	srv := Server{
		addr: fmt.Sprintf(":%s", cfg.ServerCfg.WebPort),
		mux:  gin.New(),
	}
	srv.mux.Use(middleware.LogMiddleware())
	return &srv
}

func (s *Server) ListenAndServe() error {
	rg := s.mux.Group("/api/expense/v1")
	v1.RegisterRoutes(rg)
	err := s.mux.Run(s.addr)
	if err != nil {
		return err
	}
	return nil
}
