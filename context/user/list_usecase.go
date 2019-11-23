package user

import (
	"github.com/mmuflih/go-di-arch/domain/repository"
	"github.com/mmuflih/go-httplib/httplib"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:39
**/

type ListUsecase interface {
	List(ListRequest) (error, ListResponse)
}

type ListRequest interface {
	GetQuery() string
	GetPage() int
	GetSize() int
}

type ListResponse interface{}

type listUsecase struct {
	repo repository.UserRepository
}

func (lu listUsecase) List(req ListRequest) (error, ListResponse) {
	err, lst := lu.repo.FindAll(req.GetQuery(), req.GetPage(), req.GetSize())
	if err != nil {
		return err, nil
	}

	/** restructure response */
	var resp []listResponse
	for _, u := range lst {
		res := newListResponse(u)
		resp = append(resp, res)
	}

	count := lu.repo.FindAllCount(req.GetQuery())

	return nil, httplib.NewDataPaginate(resp, count, req.GetPage(), req.GetSize())
}

func NewListUsecase(repo repository.UserRepository) ListUsecase {
	return &listUsecase{repo}
}
