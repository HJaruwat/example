package rabbitmq

import (
	"fmt"

	streadway "github.com/streadway/amqp"
)

// EmitMassageMigrate do put message to queue
func EmitMassageMigrate(body []byte) error {
	if len(body) < 1 {
		return fmt.Errorf("body not found")
	}

	conn, err := streadway.Dial(connectionString)
	defer conn.Close()
	if err != nil {
		return failOnError(err, "Failed to connect to RabbitMQ")
	}

	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		return failOnError(err, "Failed to open a channel")
	}

	q, err := ch.QueueDeclare(
		queueNameMigrate, // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		return failOnError(err, "Failed to declare an Queue")
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		streadway.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	if err != nil {
		return failOnError(err, "Failed to publish message")
	}

	return nil
}
