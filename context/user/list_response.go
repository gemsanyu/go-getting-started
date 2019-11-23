package user

import (
	"github.com/mmuflih/go-di-arch/domain/model"
	"github.com/mmuflih/go-di-arch/lib"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 22:46
**/

type listResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Password  string `json:"-"`
	Role      string `json:"role"`
	LastLogin string `json:"last_login"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func newListResponse(u *model.User) listResponse {
	return listResponse{
		u.ID, u.Email, u.Name, u.Phone, u.Password, u.Role, u.LastLogin.Format(lib.YMDHMS),
		u.CreatedAt.Format(lib.YMDHMS), u.UpdatedAt.Format(lib.YMDHMS),
	}
}
