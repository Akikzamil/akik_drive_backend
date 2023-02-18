package main

import (
	"akik_drive/config/database"
	"akik_drive/routes"
	"github.com/labstack/echo/v4"
	"akik_drive/handlers/import"
)

func main() {
	e := echo.New()
	database.ConfigDatabase()
	routes.SetupRoutes(e)
	importData.ImportDefaultData()
	e.Logger.Fatal(e.Start(":8000"))
}