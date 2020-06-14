package data

import (
	"github.com/Jammizzle/watchlist-alert/src/logging"
	valid "github.com/asaskevich/govalidator"
	"os"

	// environment variables from .env
	_ "github.com/joho/godotenv/autoload"
)

// Configuration structure defining the database configuration
type configuration struct {
	Database string `valid:"required~The email smtp server is required (MYSQL_SCHEMA)"`
	Username string `valid:"required~The email user name is required (MYSQL_USER)"`
	Password string `valid:"required~The email password is required (MYSQL_PASS)"`
}

// dataConfig holds the database configuration
var dataConfig configuration

// initConfig loads the configuration from Environment variables into databaseConfig
func init() {
	// Set the database configuration
	dataConfig = configuration{
		Username: os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL__PASS"),
		Database: os.Getenv("MYSQL_SCHEMA"),
	}

	// Assert configuration
	if ok, err := valid.ValidateStruct(dataConfig); !ok {
		logging.Fatal(err)
	}

}
