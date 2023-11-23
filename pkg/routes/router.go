package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/travierm/go-htmx/pkg/services"
)

func InitializeRoutes(e *echo.Echo) {
	services.NewTemplateRenderer(e, "./templates/*.html")

	e.GET("/", func(c echo.Context) error {
		return homePage(c, "Travier")
	})
}

func homePage(e echo.Context, name string) error {
	res := map[string]string{
		"Name":            name,
		"ContentTemplate": "home",
	}

	return e.Render(http.StatusOK, "index", res)
}
