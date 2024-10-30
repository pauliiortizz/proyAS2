package main

import (
	queues "cursos/clients_cursos"
	controllers "cursos/controllers_cursos"
	repositories "cursos/repositories_cursos"
	services "cursos/services_cursos"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Mongo
	mainRepository := repositories.NewMongo(repositories.MongoConfig{
		Host:       "mongo",
		Port:       "27017",
		Username:   "root",
		Password:   "root",
		Database:   "courses",
		Collection: "courses",
	})

	// Rabbit
	eventsQueue := queues.NewRabbit(queues.RabbitConfig{
		Host:      "localhost",
		Port:      "5672",
		Username:  "user",
		Password:  "root",
		QueueName: "courses-queue",
	})

	// Services
	service := services.NewService(mainRepository, eventsQueue)

	// Controllers
	controller := controllers.NewController(service)

	// Router
	router := gin.Default()
	router.GET("/courses/:id", controller.GetCourseByID)
	router.POST("/createCourse", controller.Create)
	router.PUT("/edit/:course_id", controller.Update)
	//router.DELETE("/hotels/:id", controller.Delete)
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("error running application: %w", err)
	}
}
