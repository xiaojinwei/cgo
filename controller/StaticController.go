package controller

import (
	"cgo/cgo"
	"net/http"
	"cgo/constant"
)

//静态资源
type StaticController struct {
	cgo.ApiController
}

func (p *StaticController)Router(router *cgo.RouterHandler)  {
	router.Router(constant.STATIC_BAES_PATH,p.img)
}

var static = http.StripPrefix(constant.STATIC_BAES_PATH, http.FileServer(http.Dir(constant.BASE_IMAGE_ADDRESS)))

func (p *StaticController)img(w http.ResponseWriter,r *http.Request)  {
	static.ServeHTTP(w,r)
}
