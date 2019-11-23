package profile

import (
	"net/http"

	"github.com/heroku/go-getting-started/context/profile"
	"github.com/heroku/go-getting-started/context/user"
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
	auth user.GetAuthUserUsecase
	pc   profile.EditUsecase
	rr   httplib.RequestReader
}

func (eh editHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := baseRequest{}
	err := eh.rr.GetJsonData(r, &req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	req.UserID = eh.auth.GetUserID(r)

	err, resp := eh.pc.Edit(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewEditHandler(auth user.GetAuthUserUsecase, pc profile.EditUsecase, rr httplib.RequestReader) EditHandler {
	return &editHandler{auth, pc, rr}
}
