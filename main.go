package main

import (
	"akik_drive/config/database"
	"akik_drive/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	database.ConfigDatabase()
	routes.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}