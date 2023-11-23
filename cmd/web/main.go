package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/travierm/go-htmx/pkg/routes"
)

func main() {
	e := echo.New()

	routes.InitializeRoutes(e)

	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Logger.Fatal(e.Start(":1323"))

}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)

	}
}
