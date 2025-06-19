package amimugi

import (
	"github.com/gorilla/mux"
	"github.com/withzeus/amimugi/adapters"
	a0handlers "github.com/withzeus/amimugi/handlers/auth0"
)

func UseAuth0Authentication(adp adapters.RouterAdapter) {
	a0 := a0handlers.NewAuth0Handler()
	a0.RegisterRoutes(adp)
}

func CreateMuxAdapter(router *mux.Router) *adapters.MuxAdapter {
	adp := adapters.NewMuxAdapter(router)
	return adp
}
