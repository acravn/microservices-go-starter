package messaging

import (
	"fmt"

	ampq "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *ampq.Connection
}

func NewRabbitMQ(uri string) (*RabbitMQ, error) {
	conn, err := ampq.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	return &RabbitMQ{conn: conn}, nil
}

func (r *RabbitMQ) Close() {
	if r.conn != nil {
		r.conn.Close()
	}
}
