package database

import (
	"context"
	"fmt"
	"pos-login/config"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() error {
	DSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := pgxpool.New(context.Background(), DSN)
	if err != nil {
		return err
	}

	DB = db
	return nil
}
