package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/", hello)
	e.GET("/user", show)
	e.POST("/user", display)

	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	var greed string = "Hello, World!"
	return c.String(http.StatusOK, greed)
}

func show(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	c.Logger().Info(u)
	return c.JSON(http.StatusOK, u)
}

func display(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	c.Logger().Info(u)
	return c.JSON(http.StatusCreated, u)
}
