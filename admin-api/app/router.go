package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()

	// Configuración personalizada de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // Permite solicitudes desde el frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Métodos HTTP permitidos
		AllowHeaders:     []string{"Authorization", "Content-Type"},           // Encabezados permitidos
		ExposeHeaders:    []string{"Content-Length"},                          // Encabezados visibles en el cliente
		AllowCredentials: true,                                                // Permitir envío de cookies o credenciales
		MaxAge:           12 * time.Hour,                                      // Tiempo que se cachea la política CORS
	}))
}

func StartRoute() {
	mapUrls()

	log.Info("Starting server on port 8004")
	router.Run(":8004")
}
