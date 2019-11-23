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

type ProfileRepository interface {
	DBConn() *sql.DB
	Save(u *model.Profile, tx *sql.Tx) error
	Update(u *model.Profile, tx *sql.Tx) error
	Find(id string) (error, *model.Profile)
	FindByUserID(userID string) (error, *model.Profile)
}
