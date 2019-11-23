package user

import (
	"net/http"

	"github.com/heroku/go-getting-started/context/user"
	"github.com/mmuflih/go-httplib/httplib"
)

type GetTokenHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type getTokenHandler struct {
	gtuc user.GetTokenUsecase
	rr   httplib.RequestReader
}

func NewGetTokenHandler(gtuc user.GetTokenUsecase, rr httplib.RequestReader) GetTokenHandler {
	return &getTokenHandler{gtuc, rr}
}

func (this getTokenHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := getTokenRequest{}
	err := this.rr.GetJsonData(r, &req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	err, resp := this.gtuc.Issue(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}
