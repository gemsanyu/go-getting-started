package user

import (
	"github.com/heroku/go-getting-started/loclib"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:51
**/

type baseRequest struct {
	ID       string
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (br baseRequest) GetID() string {
	return br.ID
}

func (br baseRequest) GetName() string {
	return br.Name
}

func (br baseRequest) GetUsername() string {
	return br.Username
}

func (br baseRequest) GetPassword() string {
	return loclib.GeneratePassword(br.Password)
}

func (br baseRequest) GetRole() string {
	return br.Role
}
