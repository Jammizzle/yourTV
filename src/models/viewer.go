package models

import "github.com/google/uuid"

type Viewer struct {
	ID         uuid.UUID `db:"viewer_id"`
	Name       string    `db:"viewer_name"`
	Email      string    `db:"viewer_email"`
	PushoverID string    `db:"pushover_id"`
}
