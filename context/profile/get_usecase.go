package profile

import (
	"github.com/heroku/go-getting-started/domain/repository"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:49
**/

type GetUsecase interface {
	Get(id string) (error, interface{})
	GetByUserID(userID string) (error, interface{})
}

type getUsecase struct {
	repo repository.ProfileRepository
}

func (gu getUsecase) Get(id string) (error, interface{}) {
	err, p := gu.repo.Find(id)
	if err != nil {
		return err, nil
	}

	return nil, p
}

func (gu getUsecase) GetByUserID(userID string) (error, interface{}) {
	err, p := gu.repo.FindByUserID(userID)
	if err != nil {
		return err, nil
	}

	return nil, p
}

func NewGetUsecase(repo repository.ProfileRepository) GetUsecase {
	return &getUsecase{repo}
}
