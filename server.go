package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// this is a sample code on https://github.com/labstack/echo
func main() {
	e := echo.New() // HTTPサーバーのハンドリング、Contextの生成、パラメータ処理、ルーティングの処理 etc

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.SetRenderer(t)

	e.GET("/", showHello)
	e.Run(standard.New(":1323"))
}

func showHello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", map[string]string{
		"world":  "World",
		"myName": "John",
	})
}

// Template difinition
type Template struct {
	templates *template.Template
}

// Render renderings template name of parameter
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
