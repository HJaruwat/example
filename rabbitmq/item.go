package rabbitmq

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	streadway "github.com/streadway/amqp"
)

type (
	// InputItemMultiple struct input message
	InputItemMultiple struct {
		Username      string    `json:"username"`
		ReferenceCode string    `json:"refcode"`
		ServerID      string    `json:"server_id"`
		Rewards       []Rewards `json:"rewards"`
		CreatedAt     time.Time `json:"created_at"`
	}

	// Rewards struct input reward info
	Rewards struct {
		ItemID   string `json:"item_id"`
		Amount   int    `json:"amount"`
		Option   string `json:"option"`
		ServerID string `json:"server_id"`
		Duration int    `json:"duration"`
		Name     string `json:"name"`
		Unit     string `json:"unit,omitempty"`
		Refcode  string `json:"refcode"`
		Username string `json:"username"`
	}
)

// EmitMassageItem do put message to queue
func EmitMassageItem(body *InputItemMultiple) error {

	if len(body.Rewards) < 1 {
		return fmt.Errorf("reward not found")
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
		queueNameItem, // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return failOnError(err, "Failed to declare an Queue")
	}

	// logic
	var errList []string
	for i, d := range body.Rewards {
		d.Refcode = fmt.Sprintf("%s-%d", body.ReferenceCode, i+1)
		d.ServerID = body.ServerID
		d.Username = body.Username
		bodyByte, err := json.Marshal(d)
		if err != nil {
			errList = append(errList, err.Error())
		}

		errPublish := ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			streadway.Publishing{
				ContentType: "text/plain",
				Body:        []byte(bodyByte),
			})

		if errPublish != nil {
			errList = append(errList, err.Error())
		}
	}

	if len(errList) > 0 {
		return fmt.Errorf(strings.Join(errList, ","))
	}

	return nil
}
