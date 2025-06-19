package adapters

import "net/http"

type RouterAdapter interface {
	GetHandler() http.Handler
	BindHandleFunc(method, path string, handler http.HandlerFunc)
}
