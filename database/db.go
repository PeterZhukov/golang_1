package database

import "context"

type Book struct {
	ID          int
	Title       string
	Year        int
	Public      bool
	PublisherID int
	Publisher   string
}

type DB interface {
	Open(c context.Context, pwd string)
	Close()
	ShowBooks() ([]Book, error)
	AddBooks([]Book) error
}
