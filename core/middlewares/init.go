package middlewares

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

// Method middleware for setting method
func Method(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		method := c.Request().Method
		customMethod := strings.ToUpper(c.QueryParam("_method"))
		switch customMethod {
		case http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut:
			method = customMethod
		}
		c.Request().Method = method
		c.SetRequest(c.Request())
		c.Set("Method", method)
		return next(c)
	})
}

// RequestID middleware for setting request id
func RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		id := uuid.Must(uuid.NewRandom()).String()
		c.Set("RequestID", id)
		c.Request().Header.Set("X-Request-ID", id)
		return next(c)
	})
}
