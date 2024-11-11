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
		Host:       "solr",    // Solr host
		Port:       "8983",    // Solr port
		Collection: "courses", // Collection name
	})

	// Rabbit
	eventsQueue := queues.NewRabbit(queues.RabbitConfig{
		Host:      "rabbitmq",
		Port:      "5672",
		Username:  "user",
		Password:  "root",
		QueueName: "courses-news",
	})

	// courses API
	cursosAPI := repositories.NewHTTP(repositories.HTTPConfig{
		Host: "cursos-api",
		Port: "8081",
	})

	// Crear instancia del servicio
	service := services.NewService(solrRepo, cursosAPI)

	// Iniciar el consumidor y pasarle el servicio
	if err := eventsQueue.StartConsumer(service); err != nil {
		log.Fatalf("Error running consumer: %v", err)
	}

	// Configurar y ejecutar el servidor web
	router := gin.Default()
	controller := controllers.NewController(service)
	router.GET("/search", controller.Search)
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
