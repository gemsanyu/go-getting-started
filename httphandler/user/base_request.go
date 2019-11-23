package user

import "github.com/mmuflih/go-di-arch/lib"

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:51
**/

type baseRequest struct {
	ID       string
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

func (br baseRequest) GetID() string {
	return br.ID
}

func (br baseRequest) GetName() string {
	return br.Name
}

func (br baseRequest) GetEmail() string {
	return br.Email
}

func (br baseRequest) GetPassword() string {
	return lib.GeneratePassword(br.Password)
}

func (br baseRequest) GetPhone() string {
	return br.Phone
}

func (br baseRequest) GetRole() string {
	return br.Role
}
