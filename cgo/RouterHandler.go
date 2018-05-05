package cgo

import (
	"net/http"
	"fmt"
)

var Router *RouterHandler = new(RouterHandler)

type RouterHandler struct {
}

var mux = make(map[string]func(http.ResponseWriter,*http.Request))

func (p *RouterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if fun, ok := mux[r.URL.Path]; ok {
		fun(w, r)
		return
	}
	http.Error(w, "error URL:"+r.URL.String(), http.StatusBadRequest)
}

func (p *RouterHandler) Router(relativePath string, handler func(http.ResponseWriter, *http.Request)) {
	mux[relativePath] = handler
}
