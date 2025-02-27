package db

import (
	"github.com/labstack/echo/v4"
)


func DatabaseMiddleware(queries *Queries) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            c.Set(DBKey, queries)
            return next(c)
        }
    }
}
