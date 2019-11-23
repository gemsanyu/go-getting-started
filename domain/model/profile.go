package model

import (
	"time"

	"github.com/heroku/go-getting-started/loclib"
)

// Profile of a user
type Profile struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
	Sex       string    `json:"sex"`
	AvatarURL string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewProfile(userID string, name string, birthDate time.Time, avatarURL string, sex string) *Profile {
	now := time.Now()
	profile := new(Profile)
	profile.ID = loclib.GenerateUUID()
	profile.Name = name
	profile.BirthDate = birthDate
	profile.AvatarURL = avatarURL
	profile.CreatedAt = now
	profile.UpdatedAt = now
	return profile
}
