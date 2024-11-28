package controller_admin

import (
	"admin-api/service_admin"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DockerController struct {
	Service *service_admin.DockerService
}

func NewDockerController(service *service_admin.DockerService) *DockerController {
	return &DockerController{Service: service}
}

func (dc *DockerController) ListContainers(c *gin.Context) {
	containers, err := dc.Service.ListContainers(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, containers)
}
