package controllers

import (
	"net/http"
	"strconv"
	"time"

	"example.com/crud/models"
	"github.com/labstack/echo/v4"
)

var books = []models.Book{
	{ID: 1, Title: "Go Tutorial", Isbn: strconv.Itoa(models.Isbn), Writer: "Anon", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func GetBooksController(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Response{http.StatusOK, "Success", books})
}

func GetBookByIDController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, b := range books {
		if b.ID == id {
			return c.JSON(http.StatusOK, models.Response{http.StatusOK, "Success", b})
		}
	}
	return c.JSON(http.StatusNotFound, models.Response{Code: http.StatusNotFound, Message: "Book not found"})
}

func AddBookController(c echo.Context) error {
	b := models.Book{}
	id := len(books) + 1
	b.ID = id
	b.Title = c.FormValue("title")
	b.Isbn = c.FormValue("isbn")
	b.Writer = c.FormValue("writer")
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()

	books = append(books, b)
	return c.JSON(http.StatusOK, models.Response{http.StatusOK, "Success", b})
}

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, b := range books {
		if b.ID == id {
			title := b.Title
			isbn := b.Isbn
			writer := b.Writer

			if c.FormValue("title") != "" {
				title = c.FormValue("title")
			}
			if c.FormValue("isbn") != "" {
				isbn = c.FormValue("isbn")
			}
			if c.FormValue("writer") != "" {
				writer = c.FormValue("writer")
			}

			b.Title = title
			b.Isbn = isbn
			b.Writer = writer
			b.UpdatedAt = time.Now()

			books[i] = b

			return c.JSON(http.StatusOK, models.Response{http.StatusOK, "Success", b})
		}
	}
	return c.JSON(http.StatusNotFound, models.Response{Code: http.StatusNotFound, Message: "Book not found"})
}

func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, b := range books {
		if b.ID == id {
			books = RemoveIndex(books, i)
			return c.JSON(http.StatusOK, models.Response{Code: http.StatusOK, Message: "Success"})
		}
	}
	return c.JSON(http.StatusNotFound, models.Response{Code: http.StatusNotFound, Message: "Book not found"})
}

func RemoveIndex(s []models.Book, index int) []models.Book {
	return append(s[:index], s[index+1:]...)
}
