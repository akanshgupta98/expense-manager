package server

import "github.com/gin-gonic/gin"

type Server struct {
	addr string
	mux  *gin.Engine
}
