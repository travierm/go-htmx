package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func main() {
	e := echo.New()

	logger := zerolog.New(os.Stdout)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
	NewTemplateRenderer(e, "public/*.html")

	logger.Log().Msg("hello world")
	e.GET("/", func(e echo.Context) error {
		return HomePage(e, "Travier")
	})

	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Logger.Fatal(e.Start(":1323"))

}

func HomePage(e echo.Context, name string) error {
	res := map[string]interface{}{
		"Name": name,
	}

	return e.Render(http.StatusOK, "index", res)
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
