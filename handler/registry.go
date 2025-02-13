package handler

import (
	"database/sql"

	"github.com/mackee/template-tanukistack/client"
)

type Registry struct {
	DB *sql.DB
}

func NewRegistry(cli *client.Client) *Registry {
	return &Registry{
		DB: cli.Database,
	}
}
