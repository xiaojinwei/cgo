package memory

import (
	"time"
	"sync"
	"container/list"
	"cgo/session"
)

var pder = &MemoryProvider{list:list.New()}

type SessionStore struct {
	sid string //session id 唯一标识
	timeAccessed time.Time //最后访问时间
	value map[interface{}]interface{} // session里面存储的值
}

func (p *SessionStore)Set(key,value interface{}) error {
	p.value[key] = value
	pder.SessionUpdate(p.sid)
	return nil
}

func (p *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(p.sid)
	if v,ok := p.value[key];ok{
		return v
	}else{
		return nil
	}
}

func (p *SessionStore)Delete(key interface{}) error {
	delete(p.value,key)
	pder.SessionUpdate(p.sid)
	return nil
}

func (p *SessionStore)SessionID() string {
	return p.sid
}

type MemoryProvider struct {
	lock sync.Mutex //锁
	sessions map[string]*list.Element //存储在内存
	list *list.List  //做GC
}

func (p *MemoryProvider)SessionInit(sid string) (session.Session,error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	v := make(map[interface{}]interface{},0)
	store := &SessionStore{sid:sid,timeAccessed:time.Now(),value: v}
	element := p.list.PushBack(store)
	p.sessions[sid] = element
	return store,nil
}

func (p *MemoryProvider)SessionRead(sid string) (session.Session,error ){
	if element,ok := pder.sessions[sid];ok{
		return element.Value.(*SessionStore),nil
	}else{
		sess,err := pder.SessionInit(sid)
		return sess,err
	}
	return nil,nil
}

func (p *MemoryProvider)SessionDestroy(sid string) error {
	if element,ok := pder.sessions[sid];ok{
		delete(pder.sessions,sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

func (p *MemoryProvider)SessionGC(maxLifeTime int64)  {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	for {
		element := pder.list.Back()
		if element == nil{
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxLifeTime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions,element.Value.(*SessionStore).sid)
		}else{
			break
		}
	}
}

func (p *MemoryProvider)SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element,ok := pder.sessions[sid];ok{
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

func init() {
	pder.sessions = make(map[string]*list.Element,0)
	session.Register("memory",pder)
}