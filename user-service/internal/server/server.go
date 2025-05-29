package server

import (
	"fmt"
	"user-service/internal/config"
	hv1 "user-service/internal/handlers/v1"
	hv2 "user-service/internal/handlers/v2"
	"user-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func New(cfg config.Config) *Server {
	srv := Server{
		mux:  gin.New(),
		addr: fmt.Sprintf(":%s", cfg.ServerCfg.WebPort),
	}
	return &srv
}

func (s *Server) ListenAndServe() error {

	s.RegisterMiddleWares()
	// version 1 APIs
	v1 := s.mux.Group("/api/user/v1")
	hv1.RegisterRoutes(v1)

	// version 2 APIs
	v2 := s.mux.Group("/api/user/v2")
	hv2.RegisterRoutes(v2)

	err := s.mux.Run(s.addr)
	if err != nil {
		return err
	}
	return nil

}

func (s *Server) RegisterMiddleWares() {
	s.mux.Use(middleware.CORS())
	s.mux.Use(middleware.Log())

}
