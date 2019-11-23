package user

import (
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:54
**/

type listRequest struct {
	Query string
	Page  int
	Size  int
}

func (lr listRequest) GetQuery() string {
	return lr.Query
}

func (lr listRequest) GetPage() int {
	return lr.Page
}

func (lr listRequest) GetSize() int {
	return lr.Size
}

func listFromRequest(rr httplib.RequestReader, r *http.Request) listRequest {
	q := rr.GetQuery(r, "q")
	p := rr.GetQueryInt(r, "page")
	s := rr.GetQueryInt(r, "size")
	if p == 0 {
		p = 1
	}
	if s == 0 {
		s = 1000
	}
	return listRequest{q, p, s}
}
