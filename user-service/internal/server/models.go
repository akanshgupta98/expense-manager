package server

import "github.com/gin-gonic/gin"

type Server struct {
	mux  *gin.Engine
	addr string
}
