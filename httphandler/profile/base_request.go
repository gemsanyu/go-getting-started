package profile

import (
	"time"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:51
**/

type baseRequest struct {
	ID        string
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatar_url"`
	BirthDate time.Time `json:"birth_date"`
	Sex       string    `json:"sex"`
}

func (br baseRequest) GetID() string {
	return br.ID
}

func (br baseRequest) GetUserID() string {
	return br.UserID
}

func (br baseRequest) GetName() string {
	return br.Name
}

func (br baseRequest) GetAvatarURL() string {
	return br.AvatarURL
}

func (br baseRequest) GetBirthDate() time.Time {
	return br.BirthDate
}

func (br baseRequest) GetSex() string {
	return br.Sex
}
