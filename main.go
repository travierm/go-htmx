package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	NewTemplateRenderer(e, "public/*.html")
	e.GET("/", func(e echo.Context) error {
		res := map[string]interface{}{
			"Name": "Wyndham",
		}

		return e.Render(http.StatusOK, "index", res)
	})

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

// func main() {
// 	e := echo.New()

// 	// Little bit of middlewares for housekeeping
// 	// e.Use(middleware.RemoveTrailingSlash())
// 	e.Use(middleware.Recover())
// 	// e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
// 	// 	rate.Limit(20),
// 	// )))

// 	// This will initiate our template renderer
// 	NewTemplateRenderer(e, "public/*.html")
// 	e.GET("/", func(e echo.Context) error {
// 		return e.Render(http.StatusOK, "index", nil)
// 	})

// 	//e.Logger.Fatal(e.Start(":4040"))
// 	e.Start(":4040")
// }
