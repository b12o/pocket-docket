package main

import (
	"time"
)

type Response struct {
	Data any `json:"data"`
}

type UpdateCounterRequest struct {
	NewVal int `json:"newVal"`
}

type User struct {
	Id           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	PasswordSalt string    `json:"password_salt"`
	CreatedAt    time.Time `json:"created"`
	UpdatedAt    time.Time `json:"updated"`
}
