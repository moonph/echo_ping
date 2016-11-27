package main

import (
	"github.com/labstack/echo"
	"github.com/moonph/echo_ping/handler"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	e.GET("/ping", handler.Ping())

	e.Start(":8000")
}
