package main

import (
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
