package queue

import (
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sbdoddi7/innoscripta/src/model"
)

func StartConsumer(ch *amqp091.Channel, queueName string, svc model.TransactionService) {
	msgs, _ := ch.Consume(queueName, "", true, false, false, false, nil)
	go func() {
		for d := range msgs {
			var msg model.TransactionMessage
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				log.Printf("bad message: %v", err)
				continue
			}
			if err := svc.ProcessTransaction(msg); err != nil {
				log.Printf("process error: %v", err)
			}
		}
	}()
}
