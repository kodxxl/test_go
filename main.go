package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/mars/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, "Hello, " + id + " !")
	})
	e.Logger.Fatal(e.Start(":1323"))
}