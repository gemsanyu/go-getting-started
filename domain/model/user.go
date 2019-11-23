package model

import (
	"time"

	"github.com/heroku/go-getting-started/loclib"
	"github.com/mmuflih/datelib"
)

type User struct {
	ID        string           `json:"id"`
	Username  string           `json:"username"`
	Name      string           `json:"name"`
	Password  string           `json:"-"`
	Role      string           `json:"role"`
	LastLogin datelib.NullTime `json:"last_login"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt datelib.NullTime `json:"deleted_at"`
}

func NewUser(username string, name string, password string, role string) *User {
	now := time.Now()
	user := new(User)
	user.ID = loclib.GenerateUUID()
	user.Username = username
	user.Name = name
	user.Password = password
	user.LastLogin.Valid = false
	user.CreatedAt = now
	user.UpdatedAt = now
	user.Role = role
	return user
}
