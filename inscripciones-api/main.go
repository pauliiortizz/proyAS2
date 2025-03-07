package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"inscripciones/config"
	controllers "inscripciones/controllers_inscripcion"
	"inscripciones/repositories_inscripcion"
	services "inscripciones/services_inscripcion"
	"log"
	"time"
)

func main() {
	// Inicializar repositorio MySQL
	mySQLRepository := repositories_inscripcion.NewMySQL(
		repositories_inscripcion.MySQLConfig{
			Host:     config.MySQLHost,
			Port:     config.MySQLPort,
			Database: config.MySQLDatabase,
			Username: config.MySQLUsername,
			Password: config.MySQLPassword,
		},
	)

	service := services.NewService(mySQLRepository)
	controller := controllers.NewController(service)

	// Configurar las rutas
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/insertinscripcion", controller.InsertInscripcion)
	router.GET("/inscripciones/user/:userID", controller.GetInscripcionByUserID)
	router.GET("/inscripciones/course/:courseID", controller.GetInscripcionByCourseID)

	// Iniciar el servidor
	if err := router.Run(":8083"); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}
}
