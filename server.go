package main

import (
	"io"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"golangsample/controller"
)

// this is a sample code on https://github.com/labstack/echo
func main() {
	e := echo.New() // HTTPサーバーのハンドリング、Contextの生成、パラメータ処理、ルーティングの処理 etc

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*/*.html")),
	}
	e.SetRenderer(t)

	e.Use(middleware.Logger()) // $ go get github.com/dgrijalva/jwt-go

	book := controller.NewBookController()
	e.GET("/books", book.List)
	e.GET("/books/:id", book.Show)
	e.GET("/books/new", book.Add)

	e.Run(standard.New(":1234"))
}

// Template definition
type Template struct {
	templates *template.Template
}

// Render renderings template name of parameter
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
