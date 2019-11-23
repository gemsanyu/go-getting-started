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
 * at: 2019-03-09 22:08
**/

type GetHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type getHandler struct {
	auth user.GetAuthUserUsecase
	pc   profile.GetUsecase
	rr   httplib.RequestReader
}

func (gh getHandler) Handle(w http.ResponseWriter, r *http.Request) {
	userID := gh.auth.GetUserID(r)
	err, resp := gh.pc.GetByUserID(userID)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewGetHandler(auth user.GetAuthUserUsecase, pc profile.GetUsecase, rr httplib.RequestReader) GetHandler {
	return &getHandler{auth, pc, rr}
}
