package profile

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

type EditUsecase interface {
	Edit(EditRequest) (error, interface{})
}

type EditRequest interface {
	GetUserID() string
	GetName() string
	GetBirthDate() time.Time
	GetSex() string
	GetAvatarURL() string
}

type editUsecase struct {
	repo repository.ProfileRepository
}

func (eu editUsecase) Edit(req EditRequest) (error, interface{}) {
	err, p := eu.repo.FindByUserID(req.GetUserID())
	if err != nil {
		return errors.New("Profile not found"), nil
	}

	tx, _ := eu.repo.DBConn().Begin()

	p.Name = req.GetName()
	p.BirthDate = req.GetBirthDate()
	p.Sex = req.GetSex()
	p.AvatarURL = req.GetAvatarURL()

	err = eu.repo.Update(p, tx)
	if err != nil {
		tx.Rollback()
		return err, nil
	}
	err = tx.Commit()
	return err, p
}

func NewEditUsecase(repo repository.ProfileRepository) EditUsecase {
	return &editUsecase{repo}
}
