package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reidx/dap/controlplane/pkg/apis"
)

type healthResponse struct {
	Status string `json:"status"`
}

type infoResponse struct {
	Service             string              `json:"service"`
	WorkspaceReadyPhase apis.WorkspacePhase `json:"workspaceReadyPhase"`
}

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, healthResponse{Status: "ok"})
	})

	router.GET("/readyz", func(c *gin.Context) {
		c.JSON(http.StatusOK, healthResponse{Status: "ready"})
	})

	router.GET("/api/v1/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, infoResponse{
			Service:             "api-server",
			WorkspaceReadyPhase: apis.WorkspacePhaseReady,
		})
	})

	return router
}
