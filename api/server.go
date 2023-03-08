package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/the-code-genin/busha-test/internal"
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
func NewServer(port int) (*Server, error) {
	// Create router
	if env, err := internal.DefaultConfig.Get("ENV"); err != nil {
		return nil, err
	} else if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	return &Server{router, port}, nil
}
