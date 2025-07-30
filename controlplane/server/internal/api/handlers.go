package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rxanders35/controlplane/server/internal/kubernetes"
)

func (s *Server) handleListPods() gin.HandlerFunc {
	return func(c *gin.Context) {
		pods, err := s.orchestrator.ListPods(c.Request.Context())
		if err != nil {
			log.Printf("Failed to list pods: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list pods"})
			return
		}
		c.JSON(http.StatusOK, pods)
	}
}

func (s *Server) handleCreateWorkspace() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req kubernetes.WorkspaceRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		}
		err := s.orchestrator.CreateWorkspace(c.Request.Context(), req)
		if err != nil {
			log.Printf("Failed to create workspace: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create workspace"})
			return
		}
	}
}

func (s *Server) handleSubmitPySparkJob() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req kubernetes.PySparkJobRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
			return
		}

		err := s.orchestrator.SubmitSparkJob(c.Request.Context(), req)
		if err != nil {
			log.Printf("Failed to create Spark Job: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to submit spark job"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "spark job '" + req.Name + "' created successfully"})
	}
}
