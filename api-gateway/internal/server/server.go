package server

import (
	"api-gateway/internal/config"
	v1 "api-gateway/internal/handlers/v1"
	"api-gateway/internal/middleware"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(cfg config.Config) *Server {
	server := Server{
		mux:       gin.New(),
		addr:      fmt.Sprintf(":%s", cfg.ServerCfg.WebPort),
		jwtSecret: cfg.JWTSecret,
	}
	c := cors.DefaultConfig()
	c.AllowCredentials = true
	c.AllowAllOrigins = true
	c.AllowHeaders = append(c.AllowHeaders, "Authorization")
	server.mux.Use(cors.New(c))
	server.mux.Use(middleware.LogMiddleware())

	return &server
}

func (s *Server) ListenAndServe() error {
	rg := s.mux.Group("/api/v1")
	v1.RegisterRoutes(rg, s.jwtSecret)

	err := s.mux.Run(s.addr)
	if err != nil {
		return err
	}
	return nil
}
