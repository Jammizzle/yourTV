package data

import (
	"github.com/Jammizzle/yourTV/src/logging"
	valid "github.com/asaskevich/govalidator"
	"os"

	// environment variables from .env
	_ "github.com/joho/godotenv/autoload"
)

// Configuration structure defining the database configuration
type configuration struct {
	Database string `valid:"required~The mysql database name is required (MYSQL_NAME)"`
	Host     string `valid:"required~The mysql database host is required (MYSQL_HOST)"`
	Port     string `valid:"required~The mysql database port is required (MYSQL_PORT)"`
	Username string `valid:"required~The mysql database user name is required (MYSQL_USER)"`
	Password string `valid:"required~The mysql database password is required (MYSQL_PASS)"`
}

// dataConfig holds the database configuration
var dataConfig configuration

// initConfig loads the configuration from Environment variables into databaseConfig
func init() {
	// Set the database configuration
	dataConfig = configuration{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Username: os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASS"),
		Database: os.Getenv("MYSQL_NAME"),
	}

	// Assert configuration
	if ok, err := valid.ValidateStruct(dataConfig); !ok {
		logging.Fatal(err)
	}

}
