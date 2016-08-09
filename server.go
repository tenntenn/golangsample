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
)

// this is a sample code on https://github.com/labstack/echo
func main() {
	e := echo.New() // HTTPサーバーのハンドリング、Contextの生成、パラメータ処理、ルーティングの処理 etc

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*/*.html")),
	}
	e.SetRenderer(t)

	e.Use(middleware.Logger()) // $ go get github.com/dgrijalva/jwt-go

	e.GET("/books/new", add)
	e.GET("/books/:id", show)

	e.Run(standard.New(":1234"))
}

func add(c echo.Context) error {
	return c.Render(http.StatusOK, "books/add", map[string]string{
		"world":  "World",
		"myName": "John",
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

	type Book struct {
		id      int
		title   string
		created []uint8
	}
	book := Book{}

	for row.Next() {
		if err := row.Scan(&book.id, &book.title, &book.created); err != nil {
			c.String(http.StatusInternalServerError, "cant execute sql: "+err.Error())
			return nil
		}
	}

	return c.Render(http.StatusOK, "books/show", map[string]string{
		"id":    strconv.Itoa(book.id),
		"title": book.title,
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
