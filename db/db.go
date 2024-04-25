package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectionDB() *sql.DB {
	dns := "user=postgres  dbname=gabriel_loja password=750317 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", dns)
	if err != nil {
		panic(err.Error())
	}

	return db
}