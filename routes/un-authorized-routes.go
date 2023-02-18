package routes

import (
	"akik_drive/handlers/auth"
	"akik_drive/handlers/file"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupUnAuthoriedRoutes(e *echo.Echo) {
	e.GET("/",Welcome);
	e.POST("/signUp",auth.SignUpUser)
	e.POST("/login",auth.Login)
	e.GET("/filestr/:id",file.GetFile)
	e.Static("/static", "assets")
}

func Welcome(c echo.Context) error{
	return c.String(http.StatusOK,"Welcome to akik drive");
}