package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/b12o/pocket-docket/data"
	"github.com/b12o/pocket-docket/types"
	"github.com/b12o/pocket-docket/utils"

	"github.com/labstack/echo/v5"
)

func RootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func CountHandler(c echo.Context) error {
	app, err := utils.GetPocketbaseInstance(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}

	if c.Request().Method == "GET" {
		countValue, err := data.GetCount(app)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to retrieve Counter value")
		}
		payload := map[string]int{"data": countValue}
		return c.JSON(http.StatusOK, payload)
	}

	if c.Request().Method == "POST" {
		var request types.UpdateCounterRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to Bind request body to struct")
		}

		if err := data.UpdateCount(app, request.NewVal); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to update counter")
		}

		return c.String(http.StatusOK, "200 OK")

	}
	return echo.NewHTTPError(http.StatusInternalServerError, "Unknown HTTP Request")
}

// --- USER Handlers ---

func CreateUserHandler(c echo.Context) error {
	app, err := utils.GetPocketbaseInstance(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}
	var newUser types.User
	if err := json.NewDecoder(c.Request().Body).Decode(&newUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user format")
	}
	newUserRecord, err := data.AddUserRecord(app, newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to add user to collection")
	}
	return c.JSON(http.StatusCreated, newUserRecord)
}

func GetUserHandler(c echo.Context) error {
	app, err := utils.GetPocketbaseInstance(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}
	userId := c.PathParam("userId")
	userRecord, err := data.GetUserRecord(app, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist")
	}
	return c.JSON(http.StatusOK, userRecord)
}

func UpdateUserHandler(c echo.Context) error {
	app, err := utils.GetPocketbaseInstance(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}
	userId := c.PathParam("userId")

	var updates map[string]any
	if err := json.NewDecoder(c.Request().Body).Decode(&updates); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to map request body")
	}

	userRecord, err := data.GetUserRecord(app, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist")
	}

	updatedUserRecord, err := data.UpdateUserRecord(app, userRecord, updates)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to update user record")
	}
	return c.JSON(http.StatusOK, updatedUserRecord)
}

func DeleteUserHandler(c echo.Context) error {
	app, err := utils.GetPocketbaseInstance(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}
	userId := c.PathParam("userId")

	userRecord, err := data.GetUserRecord(app, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist")
	}
	if err := data.DeleteUserRecord(app, userRecord); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User could not be deleted")
	}
	return c.String(http.StatusNoContent, "")
}

// --- TASK Handlers ---

func CreateTaskHandler(c echo.Context) error {
	app, err := utils.GetPocketbaseInstance(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}

	authUserId := c.Request().Header.Get("Authentication")
	if utils.IsEmptyOrWhitespace(authUserId) {
		return echo.NewHTTPError(http.StatusBadRequest, "Auth is missing")
	}

	_, err = data.GetUserRecord(app, authUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not exist")
	}

	var newTask types.Task
	newTask.CreatedBy = authUserId

	if err := data.DecodeAndValidateTask(c.Request().Body, &newTask, authUserId); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	newTaskRecord, err := data.AddTaskRecord(app, newTask)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newTaskRecord)
}

func GetTaskHandler(c echo.Context) error {
	app, err := utils.GetPocketbaseInstance(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}

	authUserId := c.Request().Header.Get("Authentication")
	if utils.IsEmptyOrWhitespace(authUserId) {
		return echo.NewHTTPError(http.StatusBadRequest, "Auth is missing")
	}

	_, err = data.GetUserRecord(app, authUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not exist")
	}

	taskId := c.PathParam("taskId")
	taskRecord, err := data.GetTaskRecord(app, taskId, authUserId)
	if err != nil {
		// either task does not exist, or was created by a different user. Return 403 for now
		return echo.NewHTTPError(http.StatusForbidden, "")
	}
	return c.JSON(http.StatusOK, taskRecord)
}

func UpdateTaskHandler(c echo.Context) error {
	app, err := utils.GetPocketbaseInstance(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Pocketbase instance is not available")
	}

	authUserId := c.Request().Header.Get("Authentication")
	if utils.IsEmptyOrWhitespace(authUserId) {
		return echo.NewHTTPError(http.StatusBadRequest, "Auth is missing")
	}

	_, err = data.GetUserRecord(app, authUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not exist")
	}

	taskId := c.PathParam("taskId")

	taskRecord, err := data.GetTaskRecord(app, taskId, authUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Task does not exist")
	}

	var updatedTask types.Task
	if err := data.DecodeAndValidateTask(c.Request().Body, &updatedTask, authUserId); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	updatedTaskRecord, err := data.UpdateTaskRecord(app, taskRecord, &updatedTask)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to update task")
	}
	return c.JSON(http.StatusOK, updatedTaskRecord)
}

func DeleteTaskHandler(c echo.Context) error {
	return nil
}
