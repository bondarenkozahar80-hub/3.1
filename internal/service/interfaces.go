package service

import (
	"context"
	"delayed-notifier/internal/model"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Storage interface {
	CreateNotification(context.Context, model.Notification) (*model.Notification, error)
	CancelNotification(context.Context, string) error
	GetNotificationById(context.Context, int) (*model.Notification, error)
	UpdateStatus(context.Context, string, string) error
}

type Cache interface {
	Get(string) (string, error)
	Set(int, interface{}) error
}

type Queue interface {
	Publish(body []byte) error
	Consume() (<-chan amqp.Delivery, error)
}

// type Sender interface {
// 	SendToTelegram(int, string) error
// }
