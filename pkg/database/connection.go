package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"microservice-blast-module/config/environment"
)

var Pool *pgxpool.Pool

// ConnectDB initializes the database connection pool.
func ConnectDB() error {
	databaseURL := environment.DatabaseDriver + "://" +
		environment.DatabaseUsername + ":" +
		environment.DatabasePassword + "@" +
		environment.DatabaseHost + ":" +
		environment.DatabasePort + "/" +
		environment.DatabaseName

	dbpool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}
	defer dbpool.Close()

	Pool = dbpool
	return nil
}
