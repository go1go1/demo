package session

import (
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

type RedisSessionMgr struct {
	addr       string
	passwd     string
	pool       *redis.Pool
	sessionMap map[string]Session
	rwLock     sync.RWMutex
}

func NewRedisSessionMgr() SessionMgr {
	sr := &RedisSessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return sr
}

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	if len(options) > 0 {
		r.passwd = options[0]
	}
	r.pool = newPool(addr, r.passwd)
	r.addr = addr
	return
}

func (r *RedisSessionMgr) GetSession(sessionId string) (session Session, err error) {
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()

	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = ErrSessionNotExist
		return
	}
	return
}

func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()

	id := uuid.New()
	sessionId := id.String()

	session = NewRedisSession(sessionId, r.pool)
	r.sessionMap[sessionId] = session
	return
}
