package router

import (
	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.Use(cors.Default())
}

func StartRoute() {
	mapUrls()
	log.Info("Starting server")
	err := router.Run(":8080")
	if err != nil {
		return
	}

}
