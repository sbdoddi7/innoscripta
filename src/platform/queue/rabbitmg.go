package queue

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func NewRabbitMQChannel() *amqp091.Channel {
	url := viper.GetString("RABBITMQ_URL")
	conn, err := amqp091.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open RabbitMQ channel: %v", err)
	}
	log.Println("Connected to RabbitMQ")
	return ch
}
