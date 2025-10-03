package db

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConfig struct {
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_NAME     string
	DB_PASSWORD string
}

type Adapter struct {
	db *pgxpool.Pool
}

func NewAdapter(d DBConfig) (*Adapter, error) {

	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", d.DB_USERNAME, d.DB_PASSWORD, d.DB_HOST, d.DB_PORT, d.DB_NAME)

	pool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}

	m, err := migrate.New(
		"file://../internal/adapters/db/migrations",
		databaseUrl+"?sslmode=disable")
	if err != nil {
		log.Println("folder isn't detect!")
		return nil, err
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No new migration to apply")
		} else {
			log.Println("Migration failed: %v", err)
			return nil, err
		}
	}

	return &Adapter{
		db: pool,
	}, nil
}
