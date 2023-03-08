package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server encapsulates a go gin router
type Server struct {
	router *gin.Engine
	port   int
}

// Fire up the server
func (s *Server) Start() error {
	return s.router.Run(fmt.Sprintf(":%d", s.port))
}

// Create a new Server
func NewServer(port int) *Server {
	router := gin.Default()

	return &Server{router, port}
}
