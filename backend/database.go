package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

func GetCount(app *pocketbase.PocketBase) (int, error) {
	record, err := app.Dao().FindRecordById("Counter", "loh30i7ry1384ep")
	if err != nil {
		// TODO http responses should be handled by handler. Return err instead
		return -1, echo.NewHTTPError(http.StatusInternalServerError, "Could not access collection 'Counter'")
	}

	value := record.GetInt("Value")
	return value, nil
}

func UpdateCount(app *pocketbase.PocketBase, newVal int) error {
	record, err := app.Dao().FindRecordById("Counter", "loh30i7ry1384ep")
	if err != nil {
		// TODO http responses should be handled by handler. Return err instead
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not access collection 'Counter'")
	}

	record.Set("Value", newVal)
	if err := app.Dao().SaveRecord(record); err != nil {
		return err
	}
	return nil
}

func AddUserRecord(app *pocketbase.PocketBase, newUser User) error {
	users, err := app.Dao().FindCollectionByNameOrId("users")

	return nil
}
