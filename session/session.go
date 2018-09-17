package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Session 接口
type Session interface {
	Set(key, value interface{}) error // 设置Session
	Get(key interface{}) interface{}  // 获取Session
	Del(key interface{}) error        // 删除Session
	SID() string                      // 当前Session ID
}

// Provider 接口
type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
	SessionUpdate(sid string) error
}

// Manager Session管理
type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

var providers = make(map[string]Provider)

// GetManager 获取Session Manager
func GetManager(providerName string, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	}

	return &Manager{
		cookieName:  cookieName,
		maxLifeTime: maxLifeTime,
		provider:    provider,
	}, nil
}

// RegisterProvider 注册Session 寄存器
func RegisterProvider(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}

	if _, p := providers[name]; p {
		panic("session: Register provider is existed")
	}

	providers[name] = provider
}

// GenerateSID 产生唯一的Session ID
func (m *Manager) GenerateSID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// SessionStart 启动Session功能
func (m *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session, err error) {
	if m == nil {
		return nil, fmt.Errorf("session: manager init failed")
	}

	m.lock.Lock()
	defer m.lock.Unlock()
	cookie, err := r.Cookie(m.cookieName)

	if err != nil || cookie.Value == "" {
		sid := m.GenerateSID()
		session, err = m.provider.SessionInit(sid)
		if err != nil {
			return nil, err
		}
		newCookie := http.Cookie{
			Name:     m.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(m.maxLifeTime),
		}
		http.SetCookie(w, &newCookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = m.provider.SessionRead(sid)
	}

	return
}

// SessionDestory 注销Session
func (m *Manager) SessionDestory(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}

	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.SessionDestroy(cookie.Value)
	expiredTime := time.Now()
	newCookie := http.Cookie{
		Name:     m.cookieName,
		Path:     "/",
		HttpOnly: true,
		Expires:  expiredTime,
		MaxAge:   -1,
	}
	http.SetCookie(w, &newCookie)
}

// SessionGC Session 垃圾回收
func (m *Manager) SessionGC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.SessionGC(m.maxLifeTime)
	time.AfterFunc(time.Duration(m.maxLifeTime)*time.Second, func() {
		m.SessionGC()
	})
}
