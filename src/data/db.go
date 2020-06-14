package data

import (
	"fmt"
	"github.com/Jammizzle/yourTV/src/logging"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

// MongoClient contains the application level session with the Mongo cluster
type MysqlClient struct {
	c *sqlx.DB
}

// ContextType is the type of constants used to identify net/http contexts
type ContextType int

// CreateConnection will instantiate a new database connection pool
func CreateConnection() (*MysqlClient, error) {
	logging.Info("Connecting to postgres server...")
	db, err := sqlx.Connect("pgx", fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dataConfig.Username, dataConfig.Password, dataConfig.Host,
		dataConfig.Port, dataConfig.Database,
	))
	if err != nil {
		return &MysqlClient{}, err
	}

	logging.Info("Connected")
	return &MysqlClient{c: db}, nil
}
