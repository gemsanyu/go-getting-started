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

type GetMyHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type getMyHandler struct {
	uc   user.GetUsecase
	auth user.GetAuthUserUsecase
}

func (gh getMyHandler) Handle(w http.ResponseWriter, r *http.Request) {
	userID := gh.auth.GetUserID(r)
	err, resp := gh.uc.Get(userID)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewGetMyHandler(uc user.GetUsecase, auth user.GetAuthUserUsecase) GetMyHandler {
	return &getMyHandler{uc, auth}
}
