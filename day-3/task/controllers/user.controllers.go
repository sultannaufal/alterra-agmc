package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"example.com/middleware/config"
	"example.com/middleware/lib/database"
	"example.com/middleware/lib/helpers"
	"example.com/middleware/middlewares"
	"example.com/middleware/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, models.Response{http.StatusOK, "Success", users})
}

func GetUserByIDController(c echo.Context) error {
	id := c.Param("id")
	user, err := database.GetUserByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, models.Response{Code: http.StatusNotFound, Message: "User not found"})
	}
	return c.JSON(http.StatusOK, models.Response{http.StatusOK, "Success", user})
}

func AddUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user.Password = hashedPassword

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, models.Response{http.StatusOK, "Success", user})
}

func GenerateTokenController(c echo.Context) error {
	var user models.User
	requestUser := new(models.User)
	if err := c.Bind(requestUser); err != nil {
		return err
	}
	if err := c.Validate(requestUser); err != nil {
		return err
	}
	result := config.DB.Where("email = ?", requestUser.Email).First(&user)
	if result.Error != nil || !helpers.CheckPasswordHash(requestUser.Password, user.Password) {
		r := models.Response{
			Code:    http.StatusUnauthorized,
			Message: "Incorrect email or password",
		}
		return c.JSON(http.StatusUnauthorized, r)
	}

	claims := middlewares.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "task-middleware",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1) * time.Hour)),
		},
		ID:        user.ID,
		UserAgent: c.Request().Header.Get("user-agent"),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(os.Getenv("APP_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, false)
	}

	return c.JSON(http.StatusOK, struct {
		Code        int    `json:"code"`
		AccessToken string `json:"access_token"`
	}{
		http.StatusOK,
		signedToken,
	})
}

func UpdateUserController(c echo.Context) error {
	auth := c.Get("authInfo").(jwt.MapClaims)
	auth_id := fmt.Sprintf("%v", auth["id"])
	id := c.Param("id")
	if auth_id != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot update other user")
	}
	result := config.DB.Model(&models.User{}).Where("id = ?", id)
	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	hashedPassword, err := helpers.HashPassword(c.FormValue("password"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	update := result.Updates(models.User{Name: c.FormValue("name"), Email: c.FormValue("email"), Password: hashedPassword})
	if err := update.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, models.Response{Code: http.StatusOK, Message: "Success"})
}

func DeleteUserController(c echo.Context) error {
	auth := c.Get("authInfo").(jwt.MapClaims)
	auth_id := fmt.Sprintf("%v", auth["id"])
	id := c.Param("id")
	if auth_id != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot update other user")
	}
	result := config.DB.Where("id = ?", id).Delete(&models.User{})
	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, models.Response{Code: http.StatusOK, Message: "Success"})
}
