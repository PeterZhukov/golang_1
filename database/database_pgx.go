package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type DbPgx struct {
	c  context.Context
	db *pgxpool.Pool
}

func (d *DbPgx) Open(c context.Context, pwd string) {

	db, err := pgxpool.New(c, "postgres://postgres:"+pwd+"@localhost:5432/golang")
	if err != nil {
		log.Fatal(err)
	}
	d.db = db
	d.c = c
}

func (d *DbPgx) Close() {
	d.db.Close()
	d.db = nil
}

func (d *DbPgx) ShowBooks() ([]Book, error) {
	rows, err := d.db.Query(d.c, "select id, title, year from books where id >= $1", 0)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		err := rows.Scan(&b.ID, &b.Title, &b.Year)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (d *DbPgx) AddBooks(books []Book) error {
	tx, err := d.db.Begin(d.c)
	if err != nil {
		return err
	}
	defer tx.Rollback(d.c)

	batch := &pgx.Batch{}
	for _, book := range books {
		batch.Queue("INSERT INTO books(title, year) values ($1, $2)", book.Title, book.Year)
	}
	res := tx.SendBatch(d.c, batch)
	err = res.Close()
	if err != nil {
		return err
	}
	return tx.Commit(d.c)
}
