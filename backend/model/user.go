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

func DecodeAndValidateLogin(requestBody io.Reader, l *Login) error {
	decoder := json.NewDecoder(requestBody)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(l); err != nil {
		return err
	}
	if util.IsEmptyOrWhitespace(l.Email) ||
		util.IsEmptyOrWhitespace(l.Password) {
		return fmt.Errorf("failed to validate login")
	}
	return nil
}

func DecodeAndValidateUser(requestBody io.Reader, u *User) error {
	decoder := json.NewDecoder(requestBody)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(u); err != nil {
		return fmt.Errorf("unable to decode task")
	}
	if util.IsEmptyOrWhitespace(u.Email) ||
		util.IsEmptyOrWhitespace(u.Username) {
		// util.IsEmptyOrWhitespace(u.PasswordHash) ||
		// util.IsEmptyOrWhitespace(u.PasswordSalt) {
		return fmt.Errorf("failed to create user")
	}
	return nil
}

func AddUserRecord(app *pocketbase.PocketBase, newUser User) (*models.Record, error) {
	users, err := app.Dao().FindCollectionByNameOrId("users")
	if err != nil {
		return nil, err
	}
	newUserRecord := models.NewRecord(users)
	form := forms.NewRecordUpsert(app, newUserRecord)
	form.LoadData(
		map[string]any{
			"username":  newUser.Username,
			"email":     newUser.Email,
			"passwword": newUser.Password,
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
