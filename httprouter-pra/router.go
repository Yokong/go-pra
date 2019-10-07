package main

import "net/http"

type Handle func(http.ResponseWriter, *http.Request, map[string]string)

func NotFound(w http.ResponseWriter, req *http.Request) {
	if req.Method != "CONNECT" {
		path := req.URL.Path
		if cp := CleanPath(path); cp != path && cp != req.Referer() {
			http.Redirect(w, req, cp, http.StatusMovedPermanently)
		}
	}

	http.NotFound(w, req)
}

type Router struct {
	node
	RedirectTrailingSlash   bool
	RedirectCaseInsensitive bool
	NotFound                http.HandlerFunc
	PanicHandler            func(http.ResponseWriter, *http.Request, interface{})
}

var _ http.Handler = New()

func New() *Router {
	return &Router{
		RedirectTrailingSlash:   true,
		RedirectCaseInsensitive: true,
		NotFound:                NotFound,
	}
}

func (r *Router) POST(path string, handle Handle) {
	r.Handle("POST", path, handle)
}

func (r *Router) Handle(method, path string, handle Handle) {
	if path[0] != '/' {
		panic("path must begin with '/'")
	}
	r.addRoute(method, path, handle)
}

func (r *Router) HandlerFunc(method, path string, handler http.HandlerFunc) {
	r.Handle(method, path,
		func(w http.ResponseWriter, req *http.Request, _ map[string]string) {
			handler(w, req)
		},
	)
}

func (r *Router) recv(w http.ResponseWriter, req *http.Request) {
	if rcv := recover(); rcv != nil {
		r.PanicHandler(w, req, rcv)
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r.PanicHandler != nil {
		defer r.recv(w, req)
	}

	path := req.URL.Path

	if handle, vars, tsr := r.getValue(req.Method, path); handle != nil {
		handle(w, req, vars)
		return
	} else if tsr && r.RedirectTrailingSlash && path != "/" {
		if path[len(path)-1] == '/' {
			path = path[:len(path)-1]
		} else {
			path = path + "/"
		}
		http.Redirect(w, req, path, http.StatusMovedPermanently)
		return
	} else if r.RedirectCaseInsensitive {
		fixedPath, found := r.findCaseInsensitivePath(req.Method, path, r.RedirectTrailingSlash)
		if found {
			http.Redirect(w, req, string(fixedPath), http.StatusMovedPermanently)
			return
		}
	}

	if r.NotFound != nil {
		r.NotFound(w, req)
	} else {
		http.NotFound(w, req)
	}
}
