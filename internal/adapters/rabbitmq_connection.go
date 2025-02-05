package adapters

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel

func InitRabbitMQ() error {
	var err error

	conn, err = amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
		return err
	}

	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Error opening RabbitMQ channel: %v", err)
		return err
	}

	_, err = ch.QueueDeclare("advice_queue", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error declaring queue: %v", err)
		return err
	}

	log.Println("RabbitMQ connected and queue declared")
	return nil
}

func PublishMessage(body []byte) error {
	if ch == nil {
		log.Println("Error: RabbitMQ channel not initialized")
		return fmt.Errorf("RabbitMQ channel not initialized")
	}

	err := ch.Publish("", "advice_queue", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})

	if err != nil {
		log.Printf("Error publishing message: %v", err)
		return err
	}

	log.Println("Message published in the queue")
	return nil
}

func GetRabbitMQChannel() *amqp.Channel {
	if ch == nil {
		log.Println("Warning: RabbitMQ channel not initialized")
		return nil
	}
	return ch
}

func CloseRabbitMQ() {
	if ch != nil {
		ch.Close()
		log.Println("RabbitMQ channel closed")
	}
	if conn != nil {
		conn.Close()
		log.Println("RabbitMQ connection closed")
	}
}
