package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTClaims struct {
	ID        uint   `json:"id"`
	UserAgent string `json:"user-agent"`
	jwt.RegisteredClaims
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("signing method invalid")
			}

			return []byte(os.Getenv("APP_KEY")), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		c.Set("authInfo", claims)

		return next(c)
	}
}
