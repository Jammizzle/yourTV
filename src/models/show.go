package models

import (
	"fmt"
	"github.com/Jammizzle/watchlist-alert/src/logging"
	"github.com/google/uuid"
	"github.com/gregdel/pushover"
	"regexp"
	"time"
)

type Show struct {
	ID      uuid.UUID `db:"show_id"`
	Name    string    `db:"name"`
	URL     string    `db:"url"`
	Pattern string    `db:"regex_pattern"`
	Regex   *regexp.Regexp
}

func (s *Show) CompileRegex() error {
	var err error
	s.Regex, err = regexp.Compile(s.Pattern)
	return err
}

type Shows []Show

func (s *Show) SendPushoverNotification(pushoverID string, epNumber int) error {
	logging.Infof("Sending to %s", pushoverID)
	app := pushover.New(modelConfig.PushoverApplicationID)
	recipient := pushover.NewRecipient(pushoverID)

	// Create the message to send
	message := &pushover.Message{
		Message:     fmt.Sprintf("Episode %d has just been released! Watch here: %s", epNumber, s.URL),
		Title:       fmt.Sprintf("New %s episode [%d]", s.Name, epNumber),
		Priority:    pushover.PriorityNormal,
		URL:         "http://yourapp.com/callback",
		URLTitle:    "Click here to acknowledge this notification",
		Timestamp:   time.Now().Unix(),
		CallbackURL: "http://yourapp.com/callback",
		Sound:       pushover.SoundCosmic,
	}

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		return err
	}
	return nil
}
