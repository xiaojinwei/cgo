package session

type Provider interface {
	SessionInit(sid string)(Session,error)
	SessionRead(sid string)(Session,error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

var provides = make(map[string]Provider)

func Register(name string,provider Provider)  {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _,dup := provides[name];dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}