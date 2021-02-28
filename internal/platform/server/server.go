package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/namku/API-HTTP-Go/internal/platform/server/handler/health"
)

type Server struct {
	engine   *gin.Engine
	httpAddr string
}

func New(host string, port uint) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("#{host}:#{port}"),
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
}
