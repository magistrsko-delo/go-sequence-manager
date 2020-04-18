package services

import (
	"github.com/streadway/amqp"
	"go-sequence-manager/Models"
	"log"
)

type RabbitMQ struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
	q    amqp.Queue
}

func (rabbitMQ *RabbitMQ) MessageToQueue(message []byte) error {
	err := rabbitMQ.Ch.Publish(
		"",
		rabbitMQ.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:     "text/plain",
			DeliveryMode:    amqp.Persistent,
			Priority:        0,
			Body:            message,
		})
	if err != nil {
		return err
	}
	return nil
}

func InitRabbitMQ() *RabbitMQ  {

	env := Models.GetEnvStruct()
	conn, err := amqp.Dial("amqp://" + env.RabbitUser + ":" + env.RabbitPassword + "@" + env.RabbitHost + ":" + env.RabbitPort)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		env.RabbitQueue, // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return &RabbitMQ{
		Conn: conn,
		Ch:   ch,
		q:    q,
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}