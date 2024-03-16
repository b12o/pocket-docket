package data

import (
	"github.com/pocketbase/pocketbase"
)

func GetCount(app *pocketbase.PocketBase) (int, error) {
	record, err := app.Dao().FindRecordById("counter", "loh30i7ry1384ep")
	if err != nil {
		return -1, err
	}

	value := record.GetInt("Value")
	return value, nil
}

func UpdateCount(app *pocketbase.PocketBase, newVal int) error {
	record, err := app.Dao().FindRecordById("counter", "loh30i7ry1384ep")
	if err != nil {
		return err
	}

	record.Set("Value", newVal)
	if err := app.Dao().SaveRecord(record); err != nil {
		return err
	}
	return nil
}
