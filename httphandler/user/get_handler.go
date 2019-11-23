package user

import (
	"net/http"

	"github.com/mmuflih/go-di-arch/context/user"
	"github.com/mmuflih/go-httplib/httplib"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 22:08
**/

type GetHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type getHandler struct {
	uc user.GetUsecase
	rr httplib.RequestReader
}

func (gh getHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := gh.rr.GetRouteParam(r, "id")
	err, resp := gh.uc.Get(id)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewGetHandler(uc user.GetUsecase, rr httplib.RequestReader) GetHandler {
	return &getHandler{uc, rr}
}
