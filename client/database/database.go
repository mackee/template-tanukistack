package database

import (
	"database/sql"
	"fmt"

	"github.com/carlmjohnson/errorx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Config struct {
	DSN string `help:"The data source name for the database" env:"DATABASE_DSN" required:""`
}

func New(cfg Config) (_db *sql.DB, err error) {
	defer errorx.Trace(&err)
	db, err := sql.Open("pgx", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}
	return db, nil
}
