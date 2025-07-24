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

	for i := 0; i < 10; i++ {
		conn, err = amqp.Dial(rabbitURL)
		if err == nil {
			ch, err = conn.Channel()
			if err == nil {
				log.Println("Connected to RabbitMQ and opened channel!")
				return ch, conn, nil
			}
			// if channel fails, close connection and retry
			conn.Close()
		}
		log.Printf("RabbitMQ not ready yet (%v). Retrying in 3s...", err)
		time.Sleep(3 * time.Second)
	}

	return nil, nil, fmt.Errorf("failed to connect to RabbitMQ after retries: %w", err)
}
