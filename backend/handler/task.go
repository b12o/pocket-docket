package handler

import (
	"net/http"

	"github.com/b12o/pocket-docket/model"
	"github.com/b12o/pocket-docket/util"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
)

func HandleCreateTask(c echo.Context) error {
	app := c.Get("app").(*pocketbase.PocketBase)

	authUserId := c.Request().Header.Get("Authentication")
	if util.IsEmptyOrWhitespace(authUserId) {
		return echo.NewHTTPError(http.StatusBadRequest, "Auth is missing")
	}

	_, err := model.GetUserRecord(app, authUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not exist")
	}

	var newTask model.Task
	newTask.CreatedBy = authUserId

	if err := model.DecodeAndValidateTask(c.Request().Body, &newTask, authUserId); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	newTaskRecord, err := model.AddTaskRecord(app, newTask)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newTaskRecord)
}

func HandleGetTasks(c echo.Context) error {
	return nil
}

func HandleGetTask(c echo.Context) error {
	app := c.Get("app").(*pocketbase.PocketBase)

	authUserId := c.Request().Header.Get("Authentication")
	if util.IsEmptyOrWhitespace(authUserId) {
		return echo.NewHTTPError(http.StatusBadRequest, "Auth is missing")
	}

	_, err := model.GetUserRecord(app, authUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not exist")
	}

	taskId := c.PathParam("taskId")
	taskRecord, err := model.GetTaskRecord(app, taskId, authUserId)
	if err != nil {
		// either task does not exist, or was created by a different user. Return 403 for now
		return echo.NewHTTPError(http.StatusForbidden, "")
	}
	return c.JSON(http.StatusOK, taskRecord)
}

func HandleUpdateTask(c echo.Context) error {
	app := c.Get("app").(*pocketbase.PocketBase)

	authUserId := c.Request().Header.Get("Authentication")
	if util.IsEmptyOrWhitespace(authUserId) {
		return echo.NewHTTPError(http.StatusBadRequest, "Auth is missing")
	}
	_, err := model.GetUserRecord(app, authUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not exist")
	}

	taskId := c.PathParam("taskId")

	taskRecord, err := model.GetTaskRecord(app, taskId, authUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Task does not exist")
	}

	var updatedTask model.Task
	if err := model.DecodeAndValidateTask(c.Request().Body, &updatedTask, authUserId); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	updatedTaskRecord, err := model.UpdateTaskRecord(app, taskRecord, &updatedTask)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to update task")
	}
	return c.JSON(http.StatusOK, updatedTaskRecord)
}

func HandleDeleteTask(c echo.Context) error {
	app := c.Get("app").(*pocketbase.PocketBase)

	authUserId := c.Request().Header.Get("Authentication")
	if util.IsEmptyOrWhitespace(authUserId) {
		return echo.NewHTTPError(http.StatusBadRequest, "Auth is missing")
	}
	_, err := model.GetUserRecord(app, authUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not exist")
	}

	taskId := c.PathParam("taskId")
	taskRecord, err := model.GetTaskRecord(app, taskId, authUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Task does not exist")
	}

	if err := model.DeleteTaskRecord(app, taskRecord); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User could not be deleted")
	}
	return c.String(http.StatusNoContent, "")
}
