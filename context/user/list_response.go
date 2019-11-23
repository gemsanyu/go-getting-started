package user

import (
	"github.com/heroku/go-getting-started/domain/model"
	"github.com/mmuflih/datelib"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 22:46
**/

type listResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Password  string `json:"-"`
	Role      string `json:"role"`
	LastLogin string `json:"last_login"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func newListResponse(u *model.User) listResponse {
	lastLogin := ""
	if u.LastLogin.Valid {
		lastLogin = u.LastLogin.Time.Format(datelib.ISO1)
	}
	return listResponse{
		u.ID, u.Username, u.Name, u.Password, u.Role, lastLogin,
		u.CreatedAt.Format(datelib.ISO1), u.UpdatedAt.Format(datelib.ISO1),
	}
}
