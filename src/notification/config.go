package notification

import (
	"github.com/Jammizzle/watchlist-alert/src/logging"
	valid "github.com/asaskevich/govalidator"
	"os"
	"strconv"

	// environment variables from .env
	_ "github.com/joho/godotenv/autoload"
)

// Configuration structure defining the database configuration
type configuration struct {
	Smtp     string `valid:"required~The email smtp server is required (EMAIL_SMTP)"`
	Username string `valid:"required~The email user name is required (EMAIL_USER)"`
	Password string `valid:"required~The email password is required (EMAIL_PASS)"`
	Port     int    `valid:"required~The email port is required (EMAIL_PORT)"`
	TLS      bool   `valid:"required~The email tls is required (EMAIL_TLS)"`
}

// notifConfig holds the notification smtp configuration
var notifConfig configuration

// initConfig loads the configuration from Environment variables into databaseConfig
func init() {
	tls, err := strconv.ParseBool(os.Getenv("EMAIL_TLS"))
	if err != nil {
		logging.Fatal(err)
	}

	port, err := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	if err != nil {
		logging.Fatal(err)
	}
	// Set the database configuration
	notifConfig = configuration{
		Username: os.Getenv("EMAIL_USER"),
		Password: os.Getenv("EMAIL_PASS"),
		Smtp:     os.Getenv("EMAIL_SMTP"),
		Port:     port,
		TLS:      tls,
	}

	// Assert configuration
	if ok, err := valid.ValidateStruct(notifConfig); !ok {
		logging.Fatal(err)
	}

}
