package rabbitmq

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Conn struct {
	Channel *amqp.Channel
}

func GetConn(rabbitURL string) (Conn, error) {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return Conn{}, err
	}

	ch, err := conn.Channel()
	return Conn{
		Channel: ch,
	}, err
}

func MakeQueue(conn Conn) amqp.Queue {
	queue, err := conn.Channel.QueueDeclare(
		"add", // queue name
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Couldn't create queue %v", err)
	}

	return queue
}

func Publish(ctx context.Context, conn Conn, q amqp.Queue, mess Message) {

	body, err := json.Marshal(mess)

	if err != nil {
		log.Printf("couldnt marshal Message  before sending %v", err)
	}

	err = conn.Channel.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})

	if err != nil {
		log.Printf("coudln't send message %v", err)
	}
}

func Consume(ctx context.Context, conn Conn, q amqp.Queue) {

	msgs, err := conn.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Printf("Failed to register a consumer %v", err)
	}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
