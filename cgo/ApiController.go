package cgo

import (
	"net/http"
	"cgo/entity"
	"cgo/constant"
)

type ApiController struct {
	Controller
}

func (p *ApiController) GetUserId(w http.ResponseWriter,r *http.Request) uint {
	user := p.GetUser(w,r)
	if user == nil {
		return 0
	}
	return user.ID
}

func (p *ApiController) GetUser(w http.ResponseWriter,r *http.Request) *entity.User {
	session := GlobalSession().SessionStart(w,r)
	if session == nil {
		return nil
	}
	key_user := session.Get(constant.KEY_USER)
	if user,ok := key_user.(*entity.User);ok{
		return user
	}
	return nil
}
