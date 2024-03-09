package main

import (
	"encoding/json"
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
	newUserRecord, err := AddUserRecord(app, newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to add user to collection")
	}
	return c.JSON(http.StatusCreated, newUserRecord)
}

func GetUserHandler(c echo.Context) error {
	app, ok := c.Get("app").(*pocketbase.PocketBase)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}
	userId := c.PathParam("userId")
	userRecord, err := GetUserRecord(app, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist")
	}
	return c.JSON(http.StatusOK, userRecord)
}

func UpdateUserHandler(c echo.Context) error {
	app, ok := c.Get("app").(*pocketbase.PocketBase)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}
	userId := c.PathParam("userId")

	var updates map[string]any
	if err := json.NewDecoder(c.Request().Body).Decode(&updates); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse update request")
	}

	userRecord, err := GetUserRecord(app, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist")
	}

	updatedUserRecord, err := UpdateUserRecord(app, userRecord, updates)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to update user record")
	}
	return c.JSON(http.StatusOK, updatedUserRecord)
}

func DeleteUserHandler(c echo.Context) error {
	app, ok := c.Get("app").(*pocketbase.PocketBase)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}
	userId := c.PathParam("userId")

	userRecord, err := GetUserRecord(app, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist")
	}
	if err := DeleteUserRecord(app, userRecord); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User could not be deleted")
	}
	return c.String(http.StatusNoContent, "")
}
