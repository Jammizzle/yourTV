package models

import "github.com/google/uuid"

type Subscription struct {
	ID            uuid.UUID `db:"subscription_id"`
	ViewerID      int       `db:"viewer_id"`
	ShowID        int       `db:"show_id"`
	EpisodeNumber int       `db:"episode_number"`
}

type Subscriber struct {
	Viewer
	Show
	SubscriptionID uuid.UUID `db:"subscription_id"`
	EpisodeNumber  int       `db:"episode_number"`
}

type Subscribers []Subscriber
