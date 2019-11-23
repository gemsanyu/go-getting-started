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
 * at: 2019-03-09 22:11
**/

type ListHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type listHandler struct {
	uc user.ListUsecase
	rr httplib.RequestReader
}

func (lh listHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := listFromRequest(lh.rr, r)
	err, resp := lh.uc.List(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponsePaginate(w, resp)
}

func NewListHandler(uc user.ListUsecase, rr httplib.RequestReader) ListHandler {
	return &listHandler{uc, rr}
}
