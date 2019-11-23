package profile

import (
	"time"

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
	GetUserID() string
	GetName() string
	GetBirthDate() time.Time
	GetSex() string
	GetAvatarURL() string
}

type addUsecase struct {
	repo repository.ProfileRepository
}

func (au addUsecase) Add(req AddRequest) (error, interface{}) {
	profile := model.NewProfile(req.GetUserID(), req.GetName(), req.GetBirthDate(),
		req.GetAvatarURL(), req.GetSex())
	tx, _ := au.repo.DBConn().Begin()
	err := au.repo.Save(profile, tx)
	if err != nil {
		tx.Rollback()
		return err, nil
	}
	err = tx.Commit()
	return err, profile
}

func NewAddUsecase(repo repository.ProfileRepository) AddUsecase {
	return &addUsecase{repo}
}
