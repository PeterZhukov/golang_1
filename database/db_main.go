package database

import (
	"context"
	"os"
)

func DbMain() {
	pwd := os.Getenv("database_password")
	c := context.Background()

	var db DB
	var api *API

	db = &DbSql{}
	db.Open(c, pwd)

	api = &API{db: db}
	api.ShowBooks()
	api.AddBook()
	api.ShowBooks()

	db.Close()

	db = &DbPgx{}
	db.Open(c, pwd)

	api = &API{db: db}
	api.ShowBooks()
	api.AddBook()
	api.ShowBooks()

	db.Close()
}
