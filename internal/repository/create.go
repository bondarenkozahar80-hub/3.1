package repository

import (
	"context"
	"delayed-notifier/internal/model"
	"fmt"
)

func (r *Repository) CreateNotification(ctx context.Context, ntfn model.Notification) (*model.Notification, error) {
	query := `INSERTO INTO notifications (id, text, status, send_at) VALUES ($1, $2, $3, $4);`

	err := r.db.QueryRowContext(
		ctx, query, ntfn.ID, ntfn.Text, ntfn.Status, ntfn.SendAt,
	).Scan(&ntfn.ID, &ntfn.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("could not scan notification from db: %w", err)
	}

	return &ntfn, nil
}
