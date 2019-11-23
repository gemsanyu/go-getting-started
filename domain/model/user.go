package model

import (
	"time"

	"github.com/mmuflih/go-di-arch/lib"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	LastLogin time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(email string, name string, phone string, password string, role string) *User {
	now := time.Now()
	user := new(User)
	user.ID = lib.GenerateUUID()
	user.Email = email
	user.Name = name
	user.Phone = phone
	user.Password = password
	user.LastLogin = now
	user.CreatedAt = now
	user.UpdatedAt = now
	user.Role = role
	return user
}
