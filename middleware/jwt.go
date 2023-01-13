package middleware

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetupJwt(e *echo.Group) {
	e.Use(echojwt.JWT([]byte("secret")))
}