package handler

import (
	"encoding/json"
	"net/http"

	"github.com/b12o/pocket-docket/model"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
)

func HandleRegisterUser(c echo.Context) error {
	app := c.Get("app").(*pocketbase.PocketBase)
	var newUser model.User
	if err := model.DecodeAndValidateUser(c.Request().Body, &newUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user format")
	}

	newUserRecord, err := model.AddUserRecord(app, newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, newUserRecord)
}

func HandleLogInUser(c echo.Context) error {
	var newLogin model.Login
	if err := model.DecodeAndValidateLogin(c.Request().Body, &newLogin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}
	return c.JSON(http.StatusOK, "")
}

func HandleGetUser(c echo.Context) error {
	app := c.Get("app").(*pocketbase.PocketBase)
	userId := c.PathParam("userId")
	userRecord, err := model.GetUserRecord(app, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist")
	}
	return c.JSON(http.StatusOK, userRecord)
}

func HandleUpdateUser(c echo.Context) error {
	app := c.Get("app").(*pocketbase.PocketBase)
	userId := c.PathParam("userId")

	var updates map[string]any
	if err := json.NewDecoder(c.Request().Body).Decode(&updates); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to map request body")
	}

	userRecord, err := model.GetUserRecord(app, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist")
	}

	updatedUserRecord, err := model.UpdateUserRecord(app, userRecord, updates)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to update user record")
	}
	return c.JSON(http.StatusOK, updatedUserRecord)
}

func HandleDeleteUser(c echo.Context) error {
	app := c.Get("app").(*pocketbase.PocketBase)
	userId := c.PathParam("userId")

	userRecord, err := model.GetUserRecord(app, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist")
	}
	if err := model.DeleteUserRecord(app, userRecord); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User could not be deleted")
	}
	return c.String(http.StatusNoContent, "")
}
