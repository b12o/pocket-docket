package model

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/b12o/pocket-docket/util"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

func DecodeAndValidateTask(requestBody io.Reader, t *Task, userId string) error {
	decoder := json.NewDecoder(requestBody)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(t); err != nil {
		return fmt.Errorf("unable to decode task")
	}

	priorities := []string{"low", "medium", "high", "critical"}
	if t.CreatedBy != userId ||
		util.IsEmptyOrWhitespace(t.Title) ||
		!util.ContainsString(priorities, t.Priority) {
		return fmt.Errorf("task failed validation")
	}
	return nil
}

func AddTaskRecord(app *pocketbase.PocketBase, newTask Task) (*models.Record, error) {
	tasks, err := app.Dao().FindCollectionByNameOrId("tasks")
	if err != nil {
		return nil, err
	}
	newTaskRecord := models.NewRecord(tasks)
	form := forms.NewRecordUpsert(app, newTaskRecord)
	form.LoadData(
		map[string]any{
			"title":       newTask.Title,
			"description": newTask.Description,
			"completed":   newTask.Completed,
			"priority":    newTask.Priority,
			"created_by":  newTask.CreatedBy,
		},
	)
	if err := form.Submit(); err != nil {
		return nil, err
	}
	return newTaskRecord, nil
}
func GetTaskRecord(app *pocketbase.PocketBase, taskId string, userId string) (*models.Record, error) {
	filterString := fmt.Sprintf("id='%v' && created_by='%v'", taskId, userId)
	records, err := app.Dao().FindFirstRecordByFilter(
		"tasks",
		filterString,
	)
	if err != nil {
		return nil, err
	}
	return records, nil
}
func UpdateTaskRecord(app *pocketbase.PocketBase, record *models.Record, task *Task) (*models.Record, error) {
	form := forms.NewRecordUpsert(app, record)
	form.LoadData(
		map[string]any{
			"title":       task.Title,
			"description": task.Description,
			"completed":   task.Completed,
			"priority":    task.Priority,
			"created_by":  task.CreatedBy,
		},
	)
	if err := form.Submit(); err != nil {
		return nil, err
	}
	return record, nil
}

func DeleteTaskRecord(app *pocketbase.PocketBase, record *models.Record) error {
	if err := app.Dao().DeleteRecord(record); err != nil {
		return err
	}
	return nil
}
