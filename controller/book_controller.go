package controller

import (
	"github.com/labstack/echo"
	"golangsample/library/resource"
	"golangsample/model"
	"net/http"
	"strconv"
)

type BookController struct{}

func NewBookController() BookController {
	return BookController{}
}

func (b *BookController) List(c echo.Context) error {
	db := resource.UseDB(c)

	row, err := db.Query("select * from books")
	if err != nil {
		c.String(http.StatusInternalServerError, "cant execute sql: "+err.Error())
		return nil
	}
	defer row.Close()

	books := []model.Book{}
	for row.Next() {
		book := model.Book{}
		if err := row.Scan(&book.ID, &book.Title, &book.Created); err != nil {
			c.String(http.StatusInternalServerError, "cant execute sql: "+err.Error())
			return nil
		}
		books = append(books, book)
	}

	return c.Render(http.StatusOK, "books/list", map[string]interface{}{
		"books": books,
	})
}

func (b *BookController) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	b := model.Book{}
	book, err := b.GetByID(c, id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.Render(http.StatusOK, "books/show", map[string]string{
		"id":    strconv.Itoa(book.ID),
		"title": book.Title,
	})
}

func (b *BookController) Add(c echo.Context) error {
	return c.Render(http.StatusOK, "books/add", map[string]string{
		"world":  "World",
		"myName": "John",
	})
}
