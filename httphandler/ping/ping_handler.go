package ping

import (
	"github.com/mmuflih/go-di-arch/context/ping"
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
)

type PingHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type pingHandler struct {
	puc ping.PingUsecase
}

func NewPingHandler(puc ping.PingUsecase) PingHandler {
	return &pingHandler{puc}
}

func (this *pingHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := pingRequest{}
	resp, err := this.puc.Ping(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}
