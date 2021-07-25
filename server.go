package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type APIError struct {
	Code    int
	Message string
}

func main() {
	e := NewRouter()
	e.Logger.Fatal(e.Start(":1323"))
}

func NewRouter() *echo.Echo {
	e := echo.New()
	// e.Logger.SetLevel(log.INFO)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/user", show)
	e.POST("/user", display)
	return e
}

func hello(c echo.Context) error {
	var greed string = "Hello, World!"
	return c.String(http.StatusOK, greed)
}

func show(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		var apierr APIError
		apierr.Code = 100
		apierr.Message = "invalid request"
		c.JSON(http.StatusBadRequest, apierr)
		return err
	}
	c.Logger().Info(u)
	return c.JSON(http.StatusOK, u)
}

func display(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		var apierr APIError
		apierr.Code = 100
		apierr.Message = "invalid request"
		c.JSON(http.StatusBadRequest, apierr)
		return err
	}
	c.Logger().Info(u)
	return c.JSON(http.StatusCreated, u)
}
