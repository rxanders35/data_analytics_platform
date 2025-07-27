package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rxanders35/controlplane/server/internal/k8s"
)

type Server struct {
	engine *gin.Engine
	k8sSvc *k8s.Service
}

func NewServer(k8sSvc *k8s.Service) *Server {
	engine := gin.Default()
	server := &Server{
		engine: engine,
		k8sSvc: k8sSvc,
	}
	server.registerRoutes()
	return server
}

func (s *Server) Start(addr string) error {
	return s.engine.Run(addr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/healthcheck", s.handleHealthCheck)
	s.engine.GET("/api/v1/pods", s.handleListPods)
	s.engine.POST("/api/v1/jobs/spark", s.handleSubmitPySparkJob)
}

func (s *Server) handleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *Server) handleListPods(c *gin.Context) {
	pods, err := s.k8sSvc.ListPods(c.Request.Context())
	if err != nil {
		log.Printf("Failed to list pods: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list pods"})
		return
	}
	c.JSON(http.StatusOK, pods)
}

func (s *Server) handleSubmitPySparkJob(c *gin.Context) {
	var req k8s.PySparkJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	err := s.k8sSvc.SubmitSparkJob(c.Request.Context(), req)
	if err != nil {
		log.Printf("Failed to create Spark Job: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to submit spark job"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "spark job '" + req.Name + "' created successfully"})
}
