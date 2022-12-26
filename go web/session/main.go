package main

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

type Session interface {
	Set(key, value interface{}) error //设置session
	Get(key interface{}) interface{}  //获取session
	Delete(key interface{}) error     //删除session
	SessionID() string                //返回sessionId
}

type Provider interface {
	// SessionInit 实现session的初始化，成功则返回新的session对象
	SessionInit(sessionId string) (Session, error)
	// SessionRead 返回由相应sessionId表示的session对象。如果不存在，则以sessionId为参数调用SessionInit()方法，创建并返回一个新的session变量
	SessionRead(sessionId string) (Session, error)
	// SessionDestroy 根据给定的sessionId删除相应的session
	SessionDestroy(sessionId string) error
	// GarbageCollector 根据maxLifeTime删除过期的session变量
	GarbageCollector(maxLifeTime int64)
}

var providers = make(map[string]Provider)

// SessionManager 把provider管理器封装一下，定义一个全局的session管理器
type SessionManager struct {
	cookieName  string     //cookie的名称
	lock        sync.Mutex //锁，保证并发时数据的安全性和一致性
	provider    Provider   //管理session
	maxLifeTime int64      //超时时间
}

func NewSessionManager(providerName, cookieName string, maxLifeTime int64) (*SessionManager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgottenimport?)", providerName)
	}
	//返回一个SessionManager对象
	return &SessionManager{
		cookieName:  cookieName,
		provider:    provider,
		maxLifeTime: maxLifeTime,
	}, nil
}

// RegisterProvider 注册方法，以便可以根据provider管理器的名称来找到其对应的provider管理器
func RegisterProvider(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, p := providers[name]; p {
		panic("session: Register provider is existed")
	}
	providers[name] = provider
}

var globalSession *SessionManager

// GetSessionId 获取sessionId
func (manager *SessionManager) GetSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// SessionBegin 根据当前请求的cookie来判断是否存在有效的session，如果不存在则创建它
func (manager *SessionManager) SessionBegin(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sessionId := manager.GetSessionId()
		session, _ = manager.provider.SessionInit(sessionId)
		cookie := http.Cookie{
			Name:     manager.cookieName,
			Value:    url.QueryEscape(sessionId),
			Path:     "/",
			MaxAge:   int(manager.maxLifeTime),
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
	} else {
		sessionId, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sessionId)
	}
	return session
}

// SessionDestroy 注销session
func (manager *SessionManager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionDestroy(cookie.Value)
	expiredTime := time.Now()
	newCookie := http.Cookie{
		Name:     manager.cookieName,
		Path:     "/",
		Expires:  expiredTime,
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, &newCookie)

}

//在启动函数中开启垃圾回收
func init() {
	go globalSession.GarbageCollector()
}
func (manager *SessionManager) GarbageCollector() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.GarbageCollector(manager.maxLifeTime)
	//使用time包中的计时器功能，它会在session超时后自动调用GarbageCollector()方法
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() {
		manager.GarbageCollector()
	})
}
func main() {
	globalSession, _ = NewSessionManager("memory", "sessionId", 3600)
}
