package service

import (
	"context"
	"delayed-notifier/internal/dto"
	"delayed-notifier/internal/model"
	"encoding/json"
)

func (s *Service) CreateNotification(ctx context.Context, notification model.Notification) (*model.Notification, error) {
	// repo
	notification.Status = "active"
	notif, err := s.storage.CreateNotification(ctx, notification)
	if err != nil {
		return nil, err
	}

	// rabbitmq
	msg := dto.NotificationDTO{
		ID: notif.ID,
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	if err := s.queue.Publish(jsonData); err != nil {
		return nil, err
	}

	// redis
	if err := s.cache.Set(notif.ID, notif.Status); err != nil {
		return nil, err
	}

	return notif, nil
}
