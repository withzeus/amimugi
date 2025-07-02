package handlers

import (
	"log"
	"net/http"

	"github.com/withzeus/amimugi/adapters"
	auth0 "github.com/withzeus/amimugi/platform"
)

type Auth0Handler struct {
	platform *auth0.Auth0
}

func NewAuth0Handler() *Auth0Handler {
	auth0, err := auth0.NewAuth0()
	if err != nil {
		log.Fatalf("Failed to initialize Auth0 Service: %v", err)
	}

	return &Auth0Handler{
		platform: auth0,
	}
}

func (h *Auth0Handler) RegisterRoutes(adp adapters.RouterAdapter) {
	adp.BindHandleFunc(http.MethodGet, "/login", h.handleLogin)
	adp.BindHandleFunc(http.MethodGet, "/test", h.handleCallback)
}

func (h *Auth0Handler) handleLogin(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(`{"message" : "Hello Login"}`))
}

func (h *Auth0Handler) handleCallback(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(`{"message" : "Hello Callback"}`))
}
