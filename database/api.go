package database

import (
	"fmt"
	"log"
)

type API struct {
	db DB
}

func (a *API) ShowBooks() {
	data, err := a.db.ShowBooks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

func (a *API) AddBook() {
	data := []Book{
		{
			Title: "Chronicles of Amber",
			Year:  1970,
		},
		{
			Title: "Dune",
			Year:  1965,
		},
	}

	err := a.db.AddBooks(data)
	if err != nil {
		log.Fatal(err)
	}
}
