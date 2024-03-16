package model

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

func AddUserRecord(app *pocketbase.PocketBase, newUser User) (*models.Record, error) {
	users, err := app.Dao().FindCollectionByNameOrId("users")
	if err != nil {
		return nil, err
	}
	newUserRecord := models.NewRecord(users)
	form := forms.NewRecordUpsert(app, newUserRecord)
	form.LoadData(
		map[string]any{
			"username":      newUser.Username,
			"email":         newUser.Email,
			"password_hash": newUser.PasswordHash,
			"password_salt": newUser.PasswordSalt,
		},
	)
	if err := form.Submit(); err != nil {
		return nil, err
	}
	return newUserRecord, nil
}

func GetUserRecord(app *pocketbase.PocketBase, userId string) (*models.Record, error) {
	record, err := app.Dao().FindRecordById("users", userId)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func UpdateUserRecord(
	app *pocketbase.PocketBase,
	record *models.Record,
	updates map[string]any) (*models.Record, error) {

	form := forms.NewRecordUpsert(app, record)
	form.LoadData(updates)
	if err := form.Submit(); err != nil {
		return nil, err
	}
	return record, nil
}

func DeleteUserRecord(app *pocketbase.PocketBase, record *models.Record) error {
	if err := app.Dao().DeleteRecord(record); err != nil {
		return err
	}
	return nil
}
