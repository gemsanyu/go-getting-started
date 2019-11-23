package user

import (
	"github.com/mmuflih/go-di-arch/domain/repository"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:49
**/

type GetUsecase interface {
	Get(id string) (error, GetResponse)
}

type GetResponse interface{}

type getUsecase struct {
	repo repository.UserRepository
}

func (lu getUsecase) Get(id string) (error, GetResponse) {
	err, lst := lu.repo.Find(id)
	if err != nil {
		return err, nil
	}

	return nil, lst
}

func NewGetUsecase(repo repository.UserRepository) GetUsecase {
	return &getUsecase{repo}
}
