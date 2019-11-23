package user

import (
	"time"

	"github.com/heroku/go-getting-started/domain/repository"
	"github.com/pkg/errors"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:37
**/

type VoidUsecase interface {
	Void(userID string) (error, interface{})
}

type voidUsecase struct {
	repo repository.UserRepository
}

func (au voidUsecase) Void(userID string) (error, interface{}) {
	err, u := au.repo.Find(userID)
	if err != nil {
		return errors.New("User not found"), nil
	}

	tx, _ := au.repo.DBConn().Begin()

	u.DeletedAt.Time = time.Now()
	u.DeletedAt.Valid = true

	err = au.repo.Update(u, tx)
	if err != nil {
		tx.Rollback()
		return err, nil
	}
	err = tx.Commit()
	return err, u
}

func NewVoidUsecase(repo repository.UserRepository) VoidUsecase {
	return &voidUsecase{repo}
}
