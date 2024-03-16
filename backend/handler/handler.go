package handler

import (
	"net/http"

	"github.com/b12o/pocket-docket/model"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
)

func RootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func CountHandler(c echo.Context) error {
	app, _ := c.Get("app").(*pocketbase.PocketBase)

	if c.Request().Method == "GET" {
		countValue, err := model.GetCount(app)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to retrieve Counter value")
		}
		payload := map[string]int{"data": countValue}
		return c.JSON(http.StatusOK, payload)
	}

	if c.Request().Method == "POST" {
		var request model.UpdateCounterRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to Bind request body to struct")
		}

		if err := model.UpdateCount(app, request.NewVal); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to update counter")
		}

		return c.String(http.StatusOK, "200 OK")

	}
	return echo.NewHTTPError(http.StatusInternalServerError, "Unknown HTTP Request")
}
