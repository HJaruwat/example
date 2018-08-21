package rabbitmq

import (
	"cabal-api/config"
	"fmt"

	streadway "github.com/streadway/amqp"
)

var (
	queueNameCash    = config.Setting.App.Amqp.NameCash
	queueNameItem    = config.Setting.App.Amqp.NameItem
	queueNameMigrate = config.Setting.App.Amqp.NameMigrate
	connectionString = config.Setting.App.Amqp.Host
)

// EmitMassageCash do put message to queue
func EmitMassageCash(body []byte) error {
	if len(body) <= 0 {
		return fmt.Errorf("body can not be empty")
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
		queueNameCash, // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
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
			Body:        []byte(body),
		})
	if err != nil {
		return failOnError(err, "Failed to publish a message")
	}

	return nil
}

func failOnError(err error, message string) error {
	return fmt.Errorf("%s : %s", message, err.Error())
}
