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
 * at: 2019-03-09 21:57
**/

type UploadAvatarHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type uploadAvatarHandler struct {
	auth user.GetAuthUserUsecase
	pc   profile.UploadAvatarUsecase
	rr   httplib.RequestReader
}

func (ah uploadAvatarHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := uploadAvatarRequest{}
	err := ah.rr.GetJsonData(r, &req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	req.UserID = ah.auth.GetUserID(r)

	err, resp := ah.pc.UploadAvatar(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewUploadAvatarHandler(auth user.GetAuthUserUsecase, pc profile.UploadAvatarUsecase, rr httplib.RequestReader) UploadAvatarHandler {
	return &uploadAvatarHandler{auth, pc, rr}
}
