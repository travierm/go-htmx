package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(e *echo.Echo) {
	NewTemplateRenderer(e, "public/index.html", "views/*.html")

	e.GET("/", func(c echo.Context) error {
		return homePage(c, "Travier")
	})
}

func homePage(e echo.Context, name string) error {
	res := map[string]interface{}{
		"Name": name,
	}

	return e.Render(http.StatusOK, "index", res)
}
