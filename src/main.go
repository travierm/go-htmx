package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// logger := zerolog.New(os.Stdout)
	// // e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	// // 	LogURI:    true,
	// // 	LogStatus: true,
	// // 	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
	// // 		logger.Info().
	// // 			Str("URI", v.URI).
	// // 			Int("status", v.Status).
	// // 			Msg("request")

	// // 		return nil
	// // 	},
	// // }))

	InitializeRoutes(e)

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
