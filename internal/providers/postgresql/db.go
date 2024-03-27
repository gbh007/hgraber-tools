package postgresql

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // драйвер для PostgreSQL
)

type Database struct {
	db *sqlx.DB
}

func Connect(ctx context.Context, dataSourceName string) (*Database, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	return &Database{db: db}, nil
}
