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
 * at: 2019-03-09 22:05
**/

type EditHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type editHandler struct {
	uc user.EditUsecase
	rr httplib.RequestReader
}

func (eh editHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := baseRequest{}
	err := eh.rr.GetJsonData(r, &req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	req.ID = eh.rr.GetRouteParam(r, "id")

	err, resp := eh.uc.Edit(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewEditHandler(uc user.EditUsecase, rr httplib.RequestReader) EditHandler {
	return &editHandler{uc, rr}
}
