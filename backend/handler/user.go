package handler

import (
	"encoding/json"
	"net/http"

	"github.com/b12o/pocket-docket/model"
	"github.com/b12o/pocket-docket/util"

	"github.com/labstack/echo/v5"
)

func HandleCreateUser(c echo.Context) error {
	app, err := util.GetPocketbaseInstance(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}
	var newUser model.User
	if err := json.NewDecoder(c.Request().Body).Decode(&newUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user format")
	}
	newUserRecord, err := model.AddUserRecord(app, newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to add user to collection")
	}
	return c.JSON(http.StatusCreated, newUserRecord)
}

func HandleGetUser(c echo.Context) error {
	app, err := util.GetPocketbaseInstance(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}
	userId := c.PathParam("userId")
	userRecord, err := model.GetUserRecord(app, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist")
	}
	return c.JSON(http.StatusOK, userRecord)
}

func HandleUpdateUser(c echo.Context) error {
	app, err := util.GetPocketbaseInstance(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}
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
	app, err := util.GetPocketbaseInstance(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}
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
