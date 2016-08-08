package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
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
	_ = id

	return c.Render(http.StatusOK, "books/show", map[string]string{
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
