package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	bookJSON = `{"title": "Mock Book", "isbn": "123456789", "writer": "Mocker"}`
)

func TestGetBookController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetBooksController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.HasPrefix(rec.Body.String(), "{\"code\":200,\"message\":\"Success\",\"data\":["))
	}
}

func TestGetBookByIDController_BookExists(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, GetBookByIDController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.HasPrefix(rec.Body.String(), "{\"code\":200,\"message\":\"Success\",\"data\":{"))
	}
}

func TestGetBookByIDController_BookNotExists(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	fmt.Println(rec.Body.String())

	if assert.NoError(t, GetBookByIDController(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestAddBookController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(bookJSON))
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, AddBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.HasPrefix(rec.Body.String(), "{\"code\":201,\"message\":\"Success\",\"data\":{"))
	}
}

func TestUpdateBookController_BookExist(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, UpdateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.HasPrefix(rec.Body.String(), "{\"code\":200,\"message\":\"Success\",\"data\":{\"id\":1,\"title\":\"Go Tutorial\",\"isbn\":\"5577006791947779410\",\"writer\":\"Anon\",\"created_at\":\""))
	}
}

func TestUpdateBookController_BookNotExist(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	if assert.NoError(t, UpdateBookController(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestDeleteBookController_BookExists(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, DeleteBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.HasPrefix(rec.Body.String(), "{\"code\":200,\"message\":\"Success\",\"data\":null}"))
	}
}

func TestDeleteBookController_BookNotExists(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	if assert.NoError(t, DeleteBookController(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}
