package main

import (
	"io"
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"golangsample/library/resource"
	"golangsample/model"
)

// this is a sample code on https://github.com/labstack/echo
func main() {
	e := echo.New() // HTTPサーバーのハンドリング、Contextの生成、パラメータ処理、ルーティングの処理 etc

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*/*.html")),
	}
	e.SetRenderer(t)

	e.Use(middleware.Logger()) // $ go get github.com/dgrijalva/jwt-go

	e.GET("/books", list)
	e.GET("/books/:id", show)
	e.GET("/books/new", add)

	e.Run(standard.New(":1234"))
}

func list(c echo.Context) error {
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

func show(c echo.Context) error {
	id := c.Param("id")

	db := resource.UseDB(c)

	row, err := db.Query("select * from books where id = ?", id)
	if err != nil {
		c.String(http.StatusInternalServerError, "cant execute sql: "+err.Error())
		return nil
	}
	defer row.Close()

	book := model.Book{}
	for row.Next() {
		if err := row.Scan(&book.ID, &book.Title, &book.Created); err != nil {
			c.String(http.StatusInternalServerError, "cant execute sql: "+err.Error())
			return nil
		}
	}

	return c.Render(http.StatusOK, "books/show", map[string]string{
		"id":    strconv.Itoa(book.ID),
		"title": book.Title,
	})
}

func add(c echo.Context) error {
	return c.Render(http.StatusOK, "books/add", map[string]string{
		"world":  "World",
		"myName": "John",
	})
}

// Template definition
type Template struct {
	templates *template.Template
}

// Render renderings template name of parameter
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
