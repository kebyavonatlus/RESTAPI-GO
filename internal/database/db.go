package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	dbConnection, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return &Database{}, fmt.Errorf("Could not connect to the database: %w", err)
	}

	return &Database{
		Client: dbConnection,
	}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}
