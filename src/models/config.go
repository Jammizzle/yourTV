package models

import (
	"github.com/Jammizzle/yourTV/src/logging"
	valid "github.com/asaskevich/govalidator"
	"os"

	// environment variables from .env
	_ "github.com/joho/godotenv/autoload"
)

// Configuration structure defining the database configuration
type configuration struct {
	PushoverApplicationID string `valid:"required~The pushover application ID is required (PUSHOVER_APP_ID)"`
}

// dataConfig holds the database configuration
var modelConfig configuration

// initConfig loads the configuration from Environment variables into databaseConfig
func init() {
	// Set the database configuration
	modelConfig = configuration{
		PushoverApplicationID: os.Getenv("PUSHOVER_APP_ID"),
	}

	// Assert configuration
	if ok, err := valid.ValidateStruct(modelConfig); !ok {
		logging.Fatal(err)
	}

}
