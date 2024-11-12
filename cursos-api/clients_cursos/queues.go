package clients_cursos

import (
	"context"
	"cursos/domain_cursos"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"time"
)

type RabbitConfig struct {
	Host      string
	Port      string
	Username  string
	Password  string
	QueueName string
}

type queueProducer struct {
	channel *amqp.Channel
	queue   amqp.Queue
}

// NewRabbit crea una nueva instancia de `queueProducer`
func NewRabbit(config RabbitConfig) *queueProducer {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Username, config.Password, config.Host, config.Port))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}

	queue, err := channel.QueueDeclare(
		config.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	return &queueProducer{channel: channel, queue: queue}
}

func (q *queueProducer) Publish(cursoNuevo domain_cursos.CourseNew) error {
	// Convertir `cursoNuevo` a JSON
	body, err := json.Marshal(cursoNuevo)
	if err != nil {
		log.Debug("Error marshaling CourseNew to JSON", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = q.channel.PublishWithContext(
		ctx,
		"",
		q.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		log.Debug("Error while publishing message", err)
		return err
	}
	log.Printf("Publicando mensaje en RabbitMQ para el curso: %v", cursoNuevo)
	return nil
}
