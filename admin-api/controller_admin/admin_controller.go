package controller_admin

import (
	"admin-api/service_admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetServices(c *gin.Context) {

	c.JSON(http.StatusOK, service_admin.GetServices(c.Request.Context()))
}
