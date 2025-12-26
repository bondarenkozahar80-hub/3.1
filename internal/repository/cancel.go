package repository

import (
	"context"
	"fmt"
)

func (r *Repository) CancelNotification(ctx context.Context, id string) error {
	query := `UPDATE notifications
	SET status = canceled
	WHERE ID = %2`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to cancel status: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not cancel notification status: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("failet to cancel status. There may not be such notification: %w", err)
	}

	return nil
}
