package user

import (
	"net/http"

	"github.com/heroku/go-getting-started/context/user"
	"github.com/mmuflih/go-httplib/httplib"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 22:08
**/

type VoidHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type voidHandler struct {
	uc user.VoidUsecase
	rr httplib.RequestReader
}

func (gh voidHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := gh.rr.GetRouteParam(r, "id")
	err, resp := gh.uc.Void(id)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewVoidHandler(uc user.VoidUsecase, rr httplib.RequestReader) VoidHandler {
	return &voidHandler{uc, rr}
}
