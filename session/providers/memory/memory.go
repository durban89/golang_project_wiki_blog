package memory

import (
	"container/list"
	"sync"
	"time"

	"wiki/session"
)

// Store 存储
type Store struct {
	sid      string                      // unique session is
	lastTime time.Time                   // last save time
	value    map[interface{}]interface{} // session value save inside
}

// Provider 寄存器
type Provider struct {
	lock     sync.RWMutex             // locker
	sessions map[string]*list.Element // map in memory
	list     *list.List               // for gc
}

var memoryProvider = &Provider{list: list.New(), sessions: make(map[string]*list.Element)}

// Set Session
func (s *Store) Set(key, value interface{}) error {
	s.value[key] = value
	memoryProvider.SessionUpdate(s.sid)
	return nil
}

// Get Session
func (s *Store) Get(key interface{}) interface{} {
	memoryProvider.SessionUpdate(s.sid)
	if v, ok := s.value[key]; ok {
		return v
	}

	return nil

}

// Del Session
func (s *Store) Del(key interface{}) error {
	delete(s.value, key)
	memoryProvider.SessionUpdate(s.sid)
	return nil
}

// SID Session ID
func (s *Store) SID() string {
	return s.sid
}

// SessionInit 一个Session
func (p *Provider) SessionInit(sid string) (session.Session, error) {
	memoryProvider.lock.Lock()
	defer memoryProvider.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	store := &Store{
		sid:      sid,
		lastTime: time.Now(),
		value:    v,
	}

	res := memoryProvider.list.PushBack(store)
	memoryProvider.sessions[sid] = res

	return store, nil
}

// SessionRead 一个Session
func (p *Provider) SessionRead(sid string) (session.Session, error) {
	if v, ok := memoryProvider.sessions[sid]; ok {
		return v.Value.(*Store), nil
	}

	store, err := memoryProvider.SessionInit(sid)
	return store, err

}

// SessionDestroy 一个Session
func (p *Provider) SessionDestroy(sid string) error {
	if v, ok := memoryProvider.sessions[sid]; ok {
		delete(memoryProvider.sessions, sid)
		memoryProvider.list.Remove(v)
		return nil
	}

	return nil
}

// SessionGC 一个Session
func (p *Provider) SessionGC(maxLifeTime int64) {
	memoryProvider.lock.Lock()
	defer memoryProvider.lock.Unlock()

	for {
		v := memoryProvider.list.Back()
		if v == nil {
			break
		}

		if v.Value.(*Store).lastTime.Unix()+maxLifeTime < time.Now().Unix() {
			memoryProvider.list.Remove(v)
			delete(memoryProvider.sessions, v.Value.(*Store).sid)
		} else {
			break
		}
	}
}

// SessionUpdate 一个Session
func (p *Provider) SessionUpdate(sid string) error {
	memoryProvider.lock.Lock()
	defer memoryProvider.lock.Unlock()

	if v, ok := memoryProvider.sessions[sid]; ok {
		v.Value.(*Store).lastTime = time.Now()
		memoryProvider.list.MoveToFront(v)
	}

	return nil
}

func init() {
	memoryProvider.sessions = make(map[string]*list.Element, 0)
	session.RegisterProvider("memory", memoryProvider)
}
