package book

import (
	"context"
	"encoding/json"
	"net/http"

	"example.com/architecture/database"
	"example.com/architecture/internal/model"
	. "example.com/architecture/pkg/util/response"
	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	json_string, _ := database.Rdb.Get(context.Background(), "book").Result()
	var json_books = []byte(json_string)

	books := make([]model.Book, 0)
	err := json.Unmarshal(json_books, &books)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, Response{Code: http.StatusOK, Message: "Success", Data: books})
}

// func GetByID(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	for _, b := range books {
// 		if b.ID == id {
// 			return c.JSON(http.StatusOK, Response{Code: http.StatusOK, Message: "Success", Data: b})
// 		}
// 	}
// 	return echo.NewHTTPError(http.StatusNotFound, "Book not found")
// }

// func Create(c echo.Context) error {
// 	b := model.Book{}
// 	id := len(books) + 1
// 	b.ID = id
// 	b.Title = c.FormValue("title")
// 	b.Isbn = c.FormValue("isbn")
// 	b.Writer = c.FormValue("writer")
// 	b.CreatedAt = time.Now()
// 	b.UpdatedAt = time.Now()

// 	books = append(books, b)
// 	return c.JSON(http.StatusOK, Response{Code: http.StatusOK, Message: "Success", Data: b})
// }

// func Update(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	for i, b := range books {
// 		if b.ID == id {
// 			title := b.Title
// 			isbn := b.Isbn
// 			writer := b.Writer

// 			if c.FormValue("title") != "" {
// 				title = c.FormValue("title")
// 			}
// 			if c.FormValue("isbn") != "" {
// 				isbn = c.FormValue("isbn")
// 			}
// 			if c.FormValue("writer") != "" {
// 				writer = c.FormValue("writer")
// 			}

// 			b.Title = title
// 			b.Isbn = isbn
// 			b.Writer = writer
// 			b.UpdatedAt = time.Now()

// 			books[i] = b

// 			return c.JSON(http.StatusOK, Response{Code: http.StatusOK, Message: "Success", Data: b})
// 		}
// 	}
// 	return c.JSON(http.StatusNotFound, ErrorResponse{Code: http.StatusNotFound, Message: "Book not found"})
// }

// func Delete(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	for i, b := range books {
// 		if b.ID == id {
// 			books = RemoveIndex(books, i)
// 			return c.JSON(http.StatusOK, Response{Code: http.StatusOK, Message: "Success"})
// 		}
// 	}
// 	return c.JSON(http.StatusNotFound, ErrorResponse{Code: http.StatusNotFound, Message: "Book not found"})
// }

// func RemoveIndex(s []model.Book, index int) []model.Book {
// 	return append(s[:index], s[index+1:]...)
// }
