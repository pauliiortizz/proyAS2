package controller_admin

import (
	"admin-api/service_admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetServices(c *gin.Context) {

	services, err := service_admin.GetServices(c.Request.Context())
	if err != nil {
		// Devolver una respuesta con el error
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Devolver los servicios en caso de éxito
	c.JSON(http.StatusOK, services)
}
