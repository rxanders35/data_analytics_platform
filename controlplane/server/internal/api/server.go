package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rxanders35/controlplane/server/internal/kubernetes"
)

type Server struct {
	engine       *gin.Engine
	orchestrator *kubernetes.Orchestrator
}

func NewServer(orchestrator *kubernetes.Orchestrator) *Server {
	engine := gin.Default()
	server := &Server{
		engine:       engine,
		orchestrator: orchestrator,
	}
	server.registerRoutes()
	return server
}

func (s *Server) Start(addr string) error {
	return s.engine.Run(addr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/healthcheck", s.handleHealthCheck())
	s.engine.GET("/api/v1/pods", s.handleListPods())
	s.engine.POST("/api/v1/jobs/spark", s.handleSubmitPySparkJob())
	s.engine.POST("/api/v1/workspaces", s.handleCreateWorkspace())
}
