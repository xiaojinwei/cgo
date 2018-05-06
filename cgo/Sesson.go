package cgo

import (
	"net/http"
	"time"
	"log"
	"cgo/utils"
)

const GSESSION  = "GSESSION"

var sessions map[string]*Session = make(map[string]*Session,0)

type Session struct {
	Cookie *http.Cookie
	Value map[string]interface{}
}

func (p *Session)Get(key string) interface{} {
	return p.Value[key]
}

func (p *Session)Set(key string,value interface{})  {
	p.Value[key] = value
}

//获取session
func Get(r *http.Request) *Session{
	cookie,err := r.Cookie(GSESSION)
	if err != nil || cookie == nil || cookie.Value == "" {
		log.Println(err)
		return nil
	}
	return sessions[cookie.Value]
}

func New(r *http.Request) *Session {
	cookie,_ := r.Cookie(GSESSION)
	var session *Session
	if cookie != nil {
		session = sessions[cookie.Value]
	}
	if session == nil {
		session = newSession()
	}
	return session
}

//生成新的session
func newSession() *Session{
	cookie := newCookie()
	session := &Session{Cookie:cookie,Value: make(map[string]interface{})}
	sessions[cookie.Value] = session
	return session
}

func newCookie() *http.Cookie {
	return &http.Cookie{
		Name:GSESSION,
		Value:newCookieValue(),
		Path:"/",
		//Domain:"/",
		Expires:time.Now().AddDate(0,1,0),
		//MaxAge:time.Now().Second() + (60 * 60 * 24 * 31),
	}
}

func newCookieValue() string {
	uuid,_ := utils.RandomUUID()
	value := uuid.String()
	return value
}