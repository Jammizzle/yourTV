package models

import "github.com/google/uuid"

type Subscription struct {
	ID            uuid.UUID `db:"subscription_id"`
	ViewerID      int       `db:"viewer_id"`
	ShowID        int       `db:"show_id"`
	EpisodeNumber int       `db:"episode_number"`
}

type Subscriber struct {
	ViewerID         uuid.UUID `db:"viewer_id"`
	ViewerName       string    `db:"viewer_name"`
	ViewerEmail      string    `db:"viewer_email"`
	ViewerPushoverID string    `db:"viewer_pushover_id"`
	ShowID           uuid.UUID `db:"show_id"`
	ShowName         int       `db:"show_name"`
	EpisodeNumber    int       `db:"episode_number"`
	SubscriptionID   uuid.UUID `db:"subscription_id"`
}

type Subscribers []Subscriber
