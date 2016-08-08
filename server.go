package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// this is a sample code on https://github.com/labstack/echo
func main() {
	e := echo.New() // HTTPサーバーのハンドリング、Contextの生成、パラメータ処理、ルーティングの処理 etc
	e.GET("/", showHello)
	e.Run(standard.New(":1323"))
}

func showHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
