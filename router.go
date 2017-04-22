package core

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

type Router struct {
	mux *mux.Router
}

func NewRouter() *Router {
	router := Router{
		mux.NewRouter(),
	}

	router.mux.StrictSlash(true)
	return &router
}

func (r *Router) addHandler(path string, callback ...func(http.ResponseWriter, *http.Request) bool) *mux.Route {
	return r.mux.HandleFunc(path, func(out http.ResponseWriter, in *http.Request) {
		for _, cb := range callback {
			if cb(out, in) {
				return
			}
		}
	})
}

func (r *Router) AddPostHandler(path string, callback ...func(http.ResponseWriter, *http.Request) bool) {
	r.addHandler(path, callback...).Methods("POST")
}

func (r *Router) AddGetHandler(path string, callback ...func(http.ResponseWriter, *http.Request) bool) {
	r.addHandler(path, callback...).Methods("GET")
}

func (r *Router) AddPutHandler(path string, callback ...func(http.ResponseWriter, *http.Request) bool) {
	r.addHandler(path, callback...).Methods("PUT")
}

func (r *Router) AddDeleteHandler(path string, callback ...func(http.ResponseWriter, *http.Request) bool) {
	r.addHandler(path, callback...).Methods("DELETE")
}

func (r *Router) AddSubRouter(prefix string, callback func(*Router)) {
	s := r.mux.PathPrefix(prefix).Subrouter()
	callback(&Router{s})
}

func (r *Router) Args(in *http.Request) map[string]string {
	return mux.Vars(in)
}

func (r *Router) AddStatic(prefix string, folder string) {
	currentModule := CurrentModule()
	r.mux.PathPrefix(prefix).HandlerFunc(func(out http.ResponseWriter, in *http.Request) {
		rpath := strings.SplitAfterN(in.RequestURI, prefix, 2)

		if len(rpath) < 2 {
			http.NotFound(out, in)
			return
		}

		http.ServeFile(out, in, filepath.Join(currentModule, folder, rpath[1]))
	})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *Router) FPrint(out io.Writer) {
	r.mux.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}

		fmt.Fprintln(out, t, route.GetHandler())
		return nil
	})
}
