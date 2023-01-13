package auth

import (
	"akik_drive/config/database"
	"akik_drive/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SignUpUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	result := database.DB.Where("email=?", user.Email).First(&user)
	if result.RowsAffected != 0 {
		return c.String(http.StatusFound, "User already exists")
	}
	database.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}
