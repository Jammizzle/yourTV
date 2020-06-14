package data

import (
	"github.com/Jammizzle/watchlist-alert/src/models"
	"github.com/google/uuid"
)

func (d *MysqlClient) GetShows() (shows models.Shows, err error) {
	tx, err := d.c.Beginx()
	if err != nil {
		return shows, err
	}

	if err := tx.Select(&shows, sqlGetShows); err != nil {
		return shows, err
	}

	return shows, nil
}

var sqlGetShows = `
	SELECT "show_id", "name", "url", "regex_pattern" FROM "Show"
`

func (d *MysqlClient) GetShowSubscribers(show uuid.UUID) (subscribers models.Subscribers, err error) {
	tx, err := d.c.Beginx()
	if err != nil {
		return subscribers, err
	}

	if err := tx.Select(&subscribers, sqlGetShowSubscribers, show); err != nil {
		return subscribers, err
	}

	return subscribers, nil
}

var sqlGetShowSubscribers = `
	SELECT "Viewer".name as viewer_name, "Viewer".email as viewer_email, "Subscription".episode_number, "Viewer".pushover_id as viewer_pushover_id 
	FROM "Viewer"
	JOIN "Subscription" ON "Subscription"."viewer_id" = "Viewer"."viewer_id"
	AND "Subscription"."show_id" = $1
`
