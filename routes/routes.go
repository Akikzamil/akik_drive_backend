package routes

import (
	"akik_drive/middleware"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	middleware.InitialiseCommonMiddlewares(e)
	e.POST("/ll", upload)
	SetupUnAuthoriedRoutes(e)
	r := e.Group("/restricted")
	middleware.SetupJwt(r)
	SetupAuthorizedRoutes(r)
}

func upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(fmt.Sprintf("assets/xx%s", file.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
}
