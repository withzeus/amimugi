package adapters

import (
	"net/http"

	"github.com/gorilla/mux"
)

type MuxAdapter struct {
	router *mux.Router
}

func NewMuxAdapter(router *mux.Router) *MuxAdapter {
	return &MuxAdapter{router: router}
}

func (mx *MuxAdapter) BindHandleFunc(method, path string, handler http.HandlerFunc) {
	mx.router.HandleFunc(path, handler).Methods(method)
}

func (mx *MuxAdapter) GetHandler() http.Handler {
	return mx.router
}
