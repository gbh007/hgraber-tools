package sqlite

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sqlx.DB
}

func Connect(ctx context.Context, dataSourceName string) (*Database, error) {
	db, err := sqlx.Open("sqlite3", dataSourceName+"?_fk=1")
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	return &Database{db: db}, nil
}
