package notification

import (
	"bytes"
	"crypto/tls"

	"html/template"

	"github.com/Jammizzle/watchlist-alert/src/logging"

	"github.com/Masterminds/sprig"
	"gopkg.in/gomail.v2"
)

// Channel for sending emails to goroutine
var sender chan *gomail.Message
var templateBasePath = "./assets/templates/"

func init() {
	sender = make(chan *gomail.Message)

	logging.Debug("Connecting to SMTP service")
	d := gomail.NewDialer(notifConfig.Smtp, notifConfig.Port, notifConfig.Username, notifConfig.Password)
	if notifConfig.TLS == true {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	logging.Debug("Starting email listener")
	go func() {
		for {
			logging.Info("Waiting for job...")
			m := <-sender
			logging.Info("Received email job")
			if err := d.DialAndSend(m); err != nil {
				logging.Error(err)
				return
			}
		}
	}()

	return
}

type Mail struct {
	Recipient   string
	Subject     string
	Content     string
	ContentType string
}

func (m Mail) Send() {
	msg := gomail.NewMessage()
	msg.SetHeader("From", notifConfig.Username)
	msg.SetHeader("To", m.Recipient)
	msg.SetHeader("Subject", m.Subject)
	msg.SetBody(m.ContentType, m.Content)

	logging.Info("Sending email job to channel")
	go func(ms *gomail.Message) {
		sender <- ms
		logging.Debug("Job sent")
	}(msg)

}

func (m Mail) RenderAndSend(name string, data interface{}) error {
	t := template.New(name).Funcs(sprig.FuncMap())
	t, err := t.ParseFiles(templateBasePath + name)
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return err
	}

	m.Content = tpl.String()
	m.Send()
	return nil
}
