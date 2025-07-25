package queue

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func NewRabbitMQChannel() (*amqp.Channel, *amqp.Connection, error) {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	var conn *amqp.Connection
	var ch *amqp.Channel
	var err error

	for range 10 {
		conn, err = amqp.Dial(rabbitURL)
		if err == nil {
			ch, err = conn.Channel()
			if err == nil {
				log.Println("Connected to RabbitMQ and opened channel!")
				return ch, conn, nil
			}
			conn.Close()
		}
		log.Printf("RabbitMQ not ready yet (%v). Retrying in 3s...", err)
		time.Sleep(3 * time.Second)
	}

	return nil, nil, fmt.Errorf("failed to connect to RabbitMQ after retries: %w", err)
}
