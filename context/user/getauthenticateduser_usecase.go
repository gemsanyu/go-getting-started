package user

import (
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
)

type GetAuthUserUsecase interface {
	GetUserID(*http.Request) string
}

type getAuthUserUsecase struct {
}

func NewGetAuthUserUsecase() GetAuthUserUsecase {
	return &getAuthUserUsecase{}
}

func (this getAuthUserUsecase) GetUserID(r *http.Request) string {
	userID, _ := httplib.ExtractClaim(r, "user_id")
	return userID.(string)
}
