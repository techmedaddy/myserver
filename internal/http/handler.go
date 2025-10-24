package http

import (
	"net/http"
	"github.com/techmedaddy/myserver/internal/service"
)

var svc = service.New()

func homeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	msg := svc.Hello(name)
	w.Write([]byte(msg))
}
