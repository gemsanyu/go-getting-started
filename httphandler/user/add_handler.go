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
 * at: 2019-03-09 21:57
**/

type AddHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type addHandler struct {
	uc user.AddUsecase
	rr httplib.RequestReader
}

func (ah addHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := baseRequest{}
	err := ah.rr.GetJsonData(r, &req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	err, resp := ah.uc.Add(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewAddHandler(uc user.AddUsecase, rr httplib.RequestReader) AddHandler {
	return &addHandler{uc, rr}
}
