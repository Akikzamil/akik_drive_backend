package auth

import (
	"akik_drive/config/database"
	"akik_drive/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Email string `json:"admin"`
	ID    uint   `json:id`
	jwt.RegisteredClaims
}

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
	t := getToken(user)

	return c.JSON(http.StatusCreated, echo.Map{"token": t})
}

func Login(c echo.Context) error {
	user := new(models.User)
	user2 := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.String(503, err.Error())
	}
	result := database.DB.Where("email=?", user.Email).First(&user2)
	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Invalid Email")
	}
	if user.Password != user2.Password{
		return c.String(http.StatusBadGateway, "Invalid Password")
	}
	t := getToken(user)

	return c.JSON(http.StatusOK, echo.Map{"token": t})
}

func getToken(user *models.User) string {
	claims := &jwtCustomClaims{Name: user.Name, Email: user.Email, ID: user.ID, RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72))}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, _ := token.SignedString([]byte("secret"))
	return t
}
