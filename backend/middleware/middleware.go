package middleware

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
)

func InjectPocketBaseAppMiddleware(app *pocketbase.PocketBase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("app", app)
			return next(c)
		}
	}
}

func VerifyPocketBaseInjectionMiddleware(app *pocketbase.PocketBase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, ok := c.Get("app").(*pocketbase.PocketBase)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, "pocketbase instance not available")
			}
			return next(c)
		}
	}
}
