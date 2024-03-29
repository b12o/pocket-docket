package util

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func RegisterRoute(
	app *pocketbase.PocketBase,
	method string,
	path string,
	handler echo.HandlerFunc) {

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		switch method {
		case "GET":
			e.Router.GET(path, handler)
		case "PUT":
			e.Router.PUT(path, handler)
		case "POST":
			e.Router.POST(path, handler)
		case "DELETE":
			e.Router.DELETE(path, handler)
		case "PATCH":
			e.Router.PATCH(path, handler)
		case "OPTIONS":
			e.Router.OPTIONS(path, handler)
		case "HEAD":
			e.Router.HEAD(path, handler)
		}
		return nil
	})
}

func IsEmptyOrWhitespace(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func GetPocketbaseInstance(c echo.Context) (*pocketbase.PocketBase, error) {
	app, ok := c.Get("app").(*pocketbase.PocketBase)
	if !ok {
		return nil, fmt.Errorf("pocketbase instance is not avaiable")
	}
	return app, nil
}

func ContainsString(list []string, s string) bool {
	for _, item := range list {
		if strings.EqualFold(item, s) {
			return true
		}
	}
	return false
}
