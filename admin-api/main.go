package admin_api

import (
	controllers "admin-api/controller_admin"
	services "admin-api/service_admin"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar el servicio Docker
	dockerService, err := services.NewDockerService()
	if err != nil {
		log.Fatalf("Error initializing Docker service: %v", err)
	}

	// Crear el controlador
	dockerController := controllers.NewDockerController(dockerService)

	// Configurar las rutas
	r := gin.Default()
	r.GET("/containers", dockerController.ListContainers)

	// Iniciar el servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
