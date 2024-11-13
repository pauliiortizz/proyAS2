package queues

import (
	"encoding/json"
	"fmt"
	sirup "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"log"
	"search/domain_search"
	"search/services_search"
)

type RabbitConfig struct {
	Host      string
	Port      string
	Username  string
	Password  string
	QueueName string
}

type Rabbit struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

func NewRabbit(config RabbitConfig) Rabbit {
	connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Username, config.Password, config.Host, config.Port))
	if err != nil {
		log.Fatalf("error getting Rabbit connection: %v", err)
	}
	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("error creating Rabbit channel: %v", err)
	}
	queue, err := channel.QueueDeclare(config.QueueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("error declaring Rabbit queue: %v", err)
	}
	return Rabbit{
		connection: connection,
		channel:    channel,
		queue:      queue,
	}
}

// StartConsumer starts listening for messages on the RabbitMQ queue using the provided service handler
func (queue Rabbit) StartConsumer(service services_search.Service) error {
	messages, err := queue.channel.Consume(
		queue.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to start consumer: %v", err)
	}

	go func() {
		for msg := range messages {
			var courseNew domain_search.CourseNew
			err := json.Unmarshal(msg.Body, &courseNew)
			if err != nil {
				sirup.Error("Error unmarshaling message:", err)
				continue
			}
			fmt.Println(courseNew)
			// Call the service method
			service.HandleCourseNew(courseNew)
		}
	}()

	return nil
}

// Close cleans up the RabbitMQ resources
func (queue Rabbit) Close() {
	if err := queue.channel.Close(); err != nil {
		sirup.Error("error closing Rabbit channel:", err)
	}
	if err := queue.connection.Close(); err != nil {
		sirup.Error("error closing Rabbit connection:", err)
	}
}
