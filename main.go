package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"finance-info-api/controller/kabutan"
	"finance-info-api/controller/mainIndex"
)

func main() {

	// make instance
	e := echo.New()

	// configure middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// config route
	e.GET("/", hello)
	e.GET("/year_high", kabutan.GetEntries())
	e.GET("/main_index/:code", mainIndex.GetEntries())

	// start the server on port number 1323
	e.Logger.Fatal(e.Start(":1323"))

}

// define handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
