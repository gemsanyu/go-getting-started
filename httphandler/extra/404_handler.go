package extra

import (
	"encoding/json"
	"net/http"
	"time"
)

type P404Handler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type p404handler struct{}

func NewP404Handler() P404Handler {
	return &p404handler{}
}

func (this p404handler) Handle(w http.ResponseWriter, r *http.Request) {
	data := r.RemoteAddr + r.RequestURI + time.Now().Format("2006-01-02:15:04:05")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(
		data,
	)
}
