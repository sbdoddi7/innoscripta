package queue

import (
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sbdoddi7/innoscripta/src/model"
)

type TransactionProducer interface {
	Publish(msg model.TransactionMessage) error
}

type transactionProducer struct {
	channel   *amqp091.Channel
	queueName string
}

func NewTransactionProducer(ch *amqp091.Channel, queueName string) TransactionProducer {
	return &transactionProducer{channel: ch, queueName: queueName}
}

func (p *transactionProducer) Publish(msg model.TransactionMessage) error {
	body, _ := json.Marshal(msg)
	return p.channel.Publish(
		"", p.queueName, false, false,
		amqp091.Publishing{ContentType: "application/json", Body: body})
}
