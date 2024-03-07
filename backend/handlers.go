package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
)

func RootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func CountHandler(c echo.Context) error {
	app, ok := c.Get("app").(*pocketbase.PocketBase)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}

	if c.Request().Method == "GET" {
		countValue, err := GetCount(app)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to retrieve Counter value")
		}
		payload := Response{
			Data: countValue,
		}
		return c.JSON(http.StatusOK, payload)
	}

	if c.Request().Method == "POST" {
		var request UpdateCounterRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to Bind request body to struct")
		}

		if err := UpdateCount(app, request.NewVal); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to update counter")
		}

		return c.String(http.StatusOK, "200 OK")

	}
	return echo.NewHTTPError(http.StatusInternalServerError, "Unknown HTTP Request")
}

func CreateUserHandler(c echo.Context) error {
	app, ok := c.Get("app").(*pocketbase.PocketBase)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}
	var newUser User
	if err := c.Bind(&newUser); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to create new user object")
	}
	if err := AddUserRecord(app, newUser); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to add user to collection")
	}
	return c.String(http.StatusCreated, "")
}

func GetUserHandler(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func UpdateUserHandler(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func DeleteUserHandler(c echo.Context) error {
	return c.String(http.StatusOK, "")
}
