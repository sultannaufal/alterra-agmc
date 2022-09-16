package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"

	"example.com/test/config"
	"example.com/test/lib/helpers"
	"example.com/test/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userForm      = `name=Mock User&email=user@mock.com&password=123`
	createdUserId uint
	isCreateUser  bool
)

func initTest() {
	config.InitDB()
	if isCreateUser {
		user := models.User{Name: "User Test", Email: "user@test.com", Password: "123"}
		hashedPassword, _ := helpers.HashPassword(user.Password)
		user.Password = hashedPassword
		if err := config.DB.Save(&user).Error; err != nil {
			return
		}
		createdUserId = user.ID
	}
}

func initEcho() *echo.Echo {
	godotenv.Load("../.env")
	initTest()
	return echo.New()
}

func TestGetUserController(t *testing.T) {
	isCreateUser = true
	e := initEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.HasPrefix(rec.Body.String(), "{\"code\":200,\"message\":\"Success\",\"data\":["))
	}
}

func TestGetUserByIDController_UserExists(t *testing.T) {
	isCreateUser = true
	e := initEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%v", createdUserId))

	if assert.NoError(t, GetUserByIDController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.HasPrefix(rec.Body.String(), "{\"code\":200,\"message\":\"Success\",\"data\":{"))
	}
}

func TestGetUserByIDController_UserNotExists(t *testing.T) {
	isCreateUser = false
	e := initEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("999999")

	if assert.NoError(t, GetUserByIDController(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestAddUserController_ValidUser(t *testing.T) {
	isCreateUser = false
	e := initEcho()
	e.Validator = helpers.Validator
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userForm))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, AddUserController(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.True(t, strings.HasPrefix(rec.Body.String(), "{\"code\":201,\"message\":\"Success\",\"data\":{"))
	}
}

func TestAddUserController_InvalidUser(t *testing.T) {
	isCreateUser = false
	e := initEcho()
	e.Validator = helpers.Validator
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("name=Mock&email=&password&"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, AddUserController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateUserController_OwnSelf(t *testing.T) {
	isCreateUser = true

	e := initEcho()
	e.Validator = helpers.Validator
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(userForm))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%v", createdUserId))
	claims := make(jwt.MapClaims)
	claims["id"] = createdUserId
	c.Set("authInfo", claims)

	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.HasPrefix(rec.Body.String(), "{\"code\":200,\"message\":\"Success\",\"data\":{\"id\":0,\"name\":\"Mock User\",\"email\":\"user@mock.com\",\"created_at\":\""))
	}
}

func TestUpdateUserController_UpdateOtherUser(t *testing.T) {
	isCreateUser = true

	e := initEcho()
	e.Validator = helpers.Validator
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(userForm))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("9999999")
	claims := make(jwt.MapClaims)
	claims["id"] = createdUserId
	c.Set("authInfo", claims)

	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}

func TestDeleteUserController_OwnSelf(t *testing.T) {
	isCreateUser = true

	e := initEcho()
	e.Validator = helpers.Validator
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%v", createdUserId))
	claims := make(jwt.MapClaims)
	claims["id"] = createdUserId
	c.Set("authInfo", claims)

	if assert.NoError(t, DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteUserController_OtherUser(t *testing.T) {
	isCreateUser = true

	e := initEcho()
	e.Validator = helpers.Validator
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("999999")
	claims := make(jwt.MapClaims)
	claims["id"] = createdUserId
	c.Set("authInfo", claims)

	if assert.NoError(t, DeleteUserController(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}
