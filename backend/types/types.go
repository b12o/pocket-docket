package types

type UpdateCounterRequest struct {
	NewVal int `json:"newVal"`
}

type User struct {
	Id           string `db:"id" json:"id"`
	Username     string `db:"username" json:"username"`
	Email        string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"password_hash"`
	PasswordSalt string `db:"password_salt" json:"password_salt"`
}

type Task struct {
	Id          string `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Completed   bool   `db:"completed" json:"completed"`
	Priority    string `db:"priority" json:"priority"`
	CreatedBy   string `db:"created_by" json:"created_by"`
}
