package repository

import (
	"context"
	"fmt"
)

func (r *Repository) UpdateStatus(ctx context.Context, id string, status string) error {
	query := `UPDATE notifications
	SET status = %1
	WHERE ID = %2`

	result, err := r.db.ExecContext(ctx, query, status, id)
	if err != nil {
		return fmt.Errorf("failed to change status: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not update notification status: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("failet to update status. There may not be such notification: %w", err)
	}

	return nil
}
