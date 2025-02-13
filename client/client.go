package client

import (
	"database/sql"

	"github.com/carlmjohnson/errorx"
	"github.com/mackee/template-tanukistack/client/database"
)

type Config struct {
	Addr     string          `help:"The address to listen on" default:":8080" env:"ADDR"`
	Database database.Config `embed:""`
}

type Client struct {
	Database *sql.DB
}

func New(cfg Config) (_cli *Client, err error) {
	defer errorx.Trace(&err)
	db, err := database.New(cfg.Database)
	if err != nil {
		return nil, err
	}

	return &Client{
		Database: db,
	}, nil
}
