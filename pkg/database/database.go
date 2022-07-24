package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	constr := "postgres://postgres:saricamou2@localhost/postgres?sslmode=disable"
	db, err := sqlx.Open("postgres", constr)
	if err != nil {
		panic(err)
	}
	return db
}
