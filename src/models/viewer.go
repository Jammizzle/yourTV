package models

import "github.com/google/uuid"

type Viewer struct {
	ID    uuid.UUID `db:"viewer_id"`
	Name  string    `db:"name"`
	Email string    `db:"email"`
}
