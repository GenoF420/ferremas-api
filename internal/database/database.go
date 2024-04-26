package database

import (
	"database/sql"
	"fmt"
	"github.com/genof420/ferremas-api/internal/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Database struct {
	*bun.DB
}

func New(cfg *config.Config) (Database, error) {
	db, err := newPostgres(cfg)
	if err != nil {
		return Database{}, err
	}

	return Database{db}, nil
}

func newPostgres(cfg *config.Config) (*bun.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)

	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	database := bun.NewDB(sqlDb, pgdialect.New())

	err := database.Ping()
	if err != nil {
		return nil, err
	}

	return database, nil
}
