package mux

import (
	"github.com/gorilla/mux"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request, context interface{})

type Handler struct {
	context interface{}
	HandlerFunc
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.HandlerFunc(w, r, h.context)
}

type Router struct {
	*mux.Router
	context interface{}
}

func (r Router) Handle(path string, h HandlerFunc) {
	r.Router.Handle(path, Handler{r.context, h})
}

func NewRouter(context interface{}) Router {
	return Router{mux.NewRouter(), context}
}
