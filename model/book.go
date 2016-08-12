package model

import (
	"github.com/labstack/echo"
	"golangsample/library/resource"
	"net/http"
)

type Book struct {
	ID      int
	Title   string
	Created []uint8
}

func (b *Book) GetByID(c echo.Context, id int) (Book, error) {
	book := Book{}

	db := resource.UseDB(c)

	row, err := db.Query("select * from books where id = ?", id)
	if err != nil {
		return book, err
	}
	defer row.Close()

	for row.Next() {
		if err := row.Scan(&book.ID, &book.Title, &book.Created); err != nil {
			return book, err
		}
	}

	return book, nil
}

func (b *Book) GetALL(c echo.Context, id int) []Book {

	db := resource.UseDB(c)

	row, err := db.Query("select * from books")
	if err != nil {
		c.String(http.StatusInternalServerError, "cant execute sql: "+err.Error())
		return nil
	}
	defer row.Close()

	books := []Book{}
	for row.Next() {
		book := Book{}
		if err := row.Scan(&book.ID, &book.Title, &book.Created); err != nil {
			c.String(http.StatusInternalServerError, "cant execute sql: "+err.Error())
			return nil
		}
		books = append(books, book)
	}

	return books
}
