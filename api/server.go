package api

import (
	"fmt"
	"strconv"

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
func NewServer(ctx *internal.AppContext) (*Server, error) {
	config, err := ctx.GetConfig()
	if err != nil {
		return nil, err
	}

	// Create router
	if env, err := config.Get("ENV"); err != nil {
		return nil, err
	} else if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	port, err := config.Get("HTTP_PORT")
	if err != nil {
		return nil, err
	}
	httpPort, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	return &Server{router, httpPort}, nil
}
