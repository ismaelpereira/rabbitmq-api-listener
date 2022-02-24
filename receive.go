package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	if err := receive(); err != nil {
		log.Fatal(err)
	}
}

func receive() error {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq/")
	if err != nil {
		return fmt.Errorf("Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"api_messages",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("Failed to declare queue", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to register a consumer", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}
