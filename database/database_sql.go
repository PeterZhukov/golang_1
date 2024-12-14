package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

type DbSql struct {
	db *sql.DB
}

func (d *DbSql) Open(c context.Context, pwd string) {
	db, err := sql.Open("pgx/v5", "postgres://postgres:"+pwd+"@127.0.0.1:5432/golang")
	if err != nil {
		log.Fatal(err)
	}
	d.db = db
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func (d *DbSql) Close() {
	d.db.Close()
	d.db = nil
}

func (d *DbSql) ShowBooks() ([]Book, error) {
	rows, err := d.db.Query("SELECT id, title, year FROM books where id > $1", 0)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []Book
	for rows.Next() {
		var b Book
		err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Year,
		)
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
func (d *DbSql) AddBooks(books []Book) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO books(title, year) values ($1, $2) RETURNING id")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, book := range books {
		res := stmt.QueryRow(book.Title, book.Year)
		var id int
		err = res.Scan(&id)
		if err != nil {
			return err
		}
		fmt.Println("Создана запись с ID:", id)
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
