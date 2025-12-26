package repository

import (
	"context"
	"database/sql"
	"delayed-notifier/internal/model"
	"fmt"
)

func (r *Repository) GetNotificationById(ctx context.Context, id int) (*model.Notification, error) {
	query := "SELECT * FROM notifications WHERE id = $1"

	var notification model.Notification
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&notification.ID,
		&notification.Text,
		&notification.Status,
		&notification.SendAt,
		&notification.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("there is not notification with such id")
		}
		return nil, fmt.Errorf("could not get notification from db: %w", err)
	}

	return &notification, nil
}
