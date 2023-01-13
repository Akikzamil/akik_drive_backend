package middleware

import (
	// "net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitialiseCommonMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 	  // Extract the credentials from HTTP request header and perform a security
	// 	  // check

	// 	  // For invalid credentials
	// 	  return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

	// 	  // For valid credentials call next
	// 	  // return next(c)
	// 	}
	//   })
	e.Static("/static", "assets")
}
