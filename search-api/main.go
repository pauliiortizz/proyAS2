package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"search/clients_search"
	controllers "search/controllers_search"
	repositories "search/repositories_search"
	services "search/services_search"
)

func main() {
	// Solr
	solrRepo := repositories.NewSolr(repositories.SolrConfig{
		Host:       "solr",   // Solr host
		Port:       "8983",   // Solr port
		Collection: "hotels", // Collection name
	})

	// Rabbit
	eventsQueue := queues.NewRabbit(queues.RabbitConfig{
		Host:      "rabbitmq",
		Port:      "5672",
		Username:  "root",
		Password:  "root",
		QueueName: "courses-news",
	})

	// Hotels API
	cursosAPI := repositories.NewHTTP(repositories.HTTPConfig{
		Host: "cursos-api",
		Port: "8081",
	})

	// Services
	service := services.NewService(solrRepo, cursosAPI)

	// Controllers
	controller := controllers.NewController(service)

	// Launch rabbit consumer
	if err := eventsQueue.StartConsumer(service.HandleCourseNew); err != nil {
		log.Fatalf("Error running consumer: %v", err)
	}

	// Create router
	router := gin.Default()
	router.GET("/search", controller.Search)
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
