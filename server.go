package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	var greed string = "Hello, World!"

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, greed)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
