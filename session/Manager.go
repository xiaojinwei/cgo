package session

import (
	"sync"
	"fmt"
	"math/rand"
	"encoding/base64"
	"net/http"
	"net/url"
	"time"
)

type Manager struct {
	cookieName string
	lock sync.Mutex
	provider Provider
	maxLifeTime int64
}

func NewManager(provideName,cookieName string,maxLifeTime int64)(*Manager,error){
	provider,ok := provides[provideName]
	if !ok {
		return nil,fmt.Errorf("session: unknown provide %q (forgotten import?)",provideName)
	}
	return &Manager{provider:provider,cookieName:cookieName,maxLifeTime:maxLifeTime},nil
}

func (p *Manager)sessionId() string {
	b := make([]byte,32)
	if _,err := rand.Read(b);err != nil{
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (p *Manager)SessionStart(w http.ResponseWriter,r *http.Request) (session Session) {
	p.lock.Lock()
	defer p.lock.Unlock()
	cookie,err := r.Cookie(p.cookieName)
	if err != nil || cookie.Value == "" {
		sid := p.sessionId()
		session,_ = p.provider.SessionInit(sid)
		cookie := http.Cookie{Name:p.cookieName,Value:url.QueryEscape(sid),Path:"/",HttpOnly:true,MaxAge:int(p.maxLifeTime)}
		http.SetCookie(w,&cookie)
	}else{
		sid,_ := url.QueryUnescape(cookie.Value)
		session,_ = p.provider.SessionRead(sid)
	}
	return
}

func (p *Manager)SessionDestroy(w http.ResponseWriter,r *http.Request)  {
	cookie,err := r.Cookie(p.cookieName)
	if err != nil || cookie.Value == ""{
		return
	}else{
		p.lock.Lock()
		defer p.lock.Unlock()
		p.provider.SessionDestroy(cookie.Value)
		cookie := http.Cookie{Name:p.cookieName,Path:"/",HttpOnly:true,Expires:time.Now(),MaxAge:-1}
		http.SetCookie(w,&cookie)
	}
}

func (p *Manager) GC() {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.provider.SessionGC(p.maxLifeTime)
	time.AfterFunc(time.Duration(p.maxLifeTime), func() {
		p.GC()
	})
}