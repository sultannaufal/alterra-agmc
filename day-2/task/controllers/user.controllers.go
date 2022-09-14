package controllers

import (
	"net/http"

	"example.com/crud/config"
	"example.com/crud/lib/database"
	"example.com/crud/models"
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

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, models.Response{http.StatusOK, "Success", user})
}

func UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	result := config.DB.Model(&models.User{}).Where("id = ?", id)
	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	update := result.Updates(models.User{Name: c.FormValue("name"), Email: c.FormValue("email"), Password: c.FormValue("password")})
	if err := update.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, models.Response{Code: http.StatusOK, Message: "Success"})
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	result := config.DB.Where("id = ?", id).Delete(models.User{})
	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, models.Response{Code: http.StatusOK, Message: "Success"})
}
