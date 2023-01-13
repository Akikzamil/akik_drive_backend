package routes

import (
	"akik_drive/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	SetupUnAuthoriedRoutes(e);
	middleware.SetupJwt(e)
	SetupAuthorizedRoutes(e)
}

