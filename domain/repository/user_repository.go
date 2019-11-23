package repository

import (
	"database/sql"

	"github.com/heroku/go-getting-started/domain/model"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:12
**/

type UserRepository interface {
	DBConn() *sql.DB
	Save(u *model.User, tx *sql.Tx) error
	Update(u *model.User, tx *sql.Tx) error
	Find(id string) (error, *model.User)
	FindByUsername(email string) (error, *model.User)
	FindAll(q string, page, size int) (error, []*model.User)
	FindAllCount(q string) int
}
