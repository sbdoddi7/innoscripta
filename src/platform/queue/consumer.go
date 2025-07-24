package queue

import (
	"encoding/json"
	"log"

	"github.com/sbdoddi7/innoscripta/src/model"
	logger "github.com/sbdoddi7/innoscripta/src/platform/log"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func StartConsumer(ch *amqp.Channel, queueName string, svc model.TransactionService) {
	msgs, _ := ch.Consume(queueName, "", true, false, false, false, nil)
	go func() {
		for d := range msgs {
			var msg model.TransactionMessage
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				log.Printf("bad message: %v", err)
				continue
			}
			logger.Logger.WithFields(logrus.Fields{
				"account_number": msg.AccountNumber,
				"amount":         msg.Amount,
				"type":           msg.Type,
			}).Info("Processing transaction event")
			if err := svc.ProcessTransaction(msg); err != nil {
				log.Printf("process error: %v", err)
			}
		}
	}()
}
