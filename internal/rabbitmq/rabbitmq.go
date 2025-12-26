package rabbitmq

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	retries = 3
	pause   = 1 * time.Second
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	Queue   amqp.Queue
}

func NewRabbitMQ(host, queueName string) (*RabbitMQ, error) {
	conn, err := connectWithRetries(host, retries, pause)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	return &RabbitMQ{
		conn:    conn,
		channel: ch,
		Queue:   q,
	}, nil
}

func (r *RabbitMQ) Publish(body []byte) error {
	err := r.channel.Publish(
		"",
		r.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	return nil
}

func (r *RabbitMQ) Consume() (<-chan amqp.Delivery, error) {
	msgs, err := r.channel.Consume(
		r.Queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to start consuming: %w", err)
	}
	return msgs, nil
}

func (r *RabbitMQ) Close() {
	if err := r.channel.Close(); err != nil {
		log.Println("failed to close channel:", err)
	}
	if err := r.conn.Close(); err != nil {
		log.Println("failed to close connection:", err)
	}
}

func connectWithRetries(url string, retries int, pause time.Duration) (*amqp.Connection, error) {
	var conn *amqp.Connection
	var err error

	for i := 0; i < retries; i++ {
		conn, err = amqp.Dial(url)
		if err == nil {
			return conn, nil
		}
		log.Printf("RabbitMQ connection failed, attempt %d/%d: %v\n", i+1, retries, err)
		time.Sleep(pause)
	}

	return nil, fmt.Errorf("failed to connect after %d attempts: %w", retries, err)
}
