package user

import (
	"github.com/heroku/go-getting-started/domain/model"
	"github.com/heroku/go-getting-started/domain/repository"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:29
**/

type AddUsecase interface {
	Add(AddRequest) (error, interface{})
}

type AddRequest interface {
	GetName() string
	GetUsername() string
	GetPassword() string
	GetRole() string
}

type addUsecase struct {
	repo repository.UserRepository
}

func (au addUsecase) Add(req AddRequest) (error, interface{}) {
	u := model.NewUser(req.GetUsername(), req.GetName(), req.GetPassword(),
		req.GetRole())

	tx, _ := au.repo.DBConn().Begin()
	err := au.repo.Save(u, tx)
	if err != nil {
		tx.Rollback()
		return err, nil
	}
	err = tx.Commit()
	return err, u
}

func NewAddUsecase(repo repository.UserRepository) AddUsecase {
	return &addUsecase{repo}
}
