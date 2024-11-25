package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Repository *Queries

func MustInitDB() *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), "postgresql://postgres@localhost:5432/tech-support")
	if err != nil {
		panic(err)
	}

	Repository = New(db)
	return db
}
