package data

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

// MongoClient contains the application level session with the Mongo cluster
type MysqlClient struct {
	c *sqlx.DB
}

// ContextType is the type of constants used to identify net/http contexts
type ContextType int

// Context WithValue constants
const (
	ContextNone               ContextType = iota // 0 - public routes
	ContextDatabaseConnection                    // 1 - Database Connection
)

// CreateConnection will instantiate a new database connection pool
func CreateConnection() (*MysqlClient, error) {
	fmt.Print("Connecting to postgres server...")

	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("pgx", "postgres://root:password@127.0.0.1:5432/helperv2")
	if err != nil {
		return &MysqlClient{}, err
	}

	fmt.Println("Connected")
	return &MysqlClient{c: db}, nil
}
