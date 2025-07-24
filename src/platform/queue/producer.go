package queue

import (
	"encoding/json"

	"github.com/sbdoddi7/innoscripta/src/model"
	logger "github.com/sbdoddi7/innoscripta/src/platform/log"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type TransactionProducer interface {
	Publish(msg model.TransactionMessage) error
}

type transactionProducer struct {
	channel   *amqp.Channel
	queueName string
}

func NewTransactionProducer(ch *amqp.Channel, queueName string) TransactionProducer {
	return &transactionProducer{channel: ch, queueName: queueName}
}

func (p *transactionProducer) Publish(msg model.TransactionMessage) error {
	logger.Logger.WithFields(logrus.Fields{
		"account_number": msg.AccountNumber,
		"amount":         msg.Amount,
		"type":           msg.Type,
	}).Info("Publishing transaction event")
	body, _ := json.Marshal(msg)
	return p.channel.Publish(
		"", p.queueName, false, false,
		amqp.Publishing{ContentType: "application/json", Body: body})
}
