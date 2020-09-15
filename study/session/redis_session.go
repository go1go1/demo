package session

import (
	"encoding/json"
	"sync"

	"github.com/gomodule/redigo/redis"
)

const (
	SessionFlagNone = iota
	SessionFlagModify
)

type RedisSession struct {
	sessionId string
	pool      *redis.Pool
	data      map[string]interface{}
	rwLock    sync.RWMutex
	flag      int
}

func NewRedisSession(sessionId string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		sessionId: sessionId,
		data:      make(map[string]interface{}, 8),
		flag:      SessionFlagNone,
		pool:      pool,
	}
	return s
}

func (r *RedisSession) loadFromRedis() (err error) {
	conn := r.pool.Get()
	reply, err := conn.Do("GET", r.sessionId)
	if err != nil {
		return
	}
	data, err := redis.String(reply, err)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(data), &r.data)
	if err != nil {
		return
	}
	return
}

func (r *RedisSession) Set(key string, value interface{}) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()

	r.data[key] = value
	r.flag = SessionFlagModify
	return
}
func (r *RedisSession) Get(key string) (result interface{}, err error) {
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()
	//延迟加载
	if r.flag == SessionFlagNone { //Session未加载
		err = r.loadFromRedis()
		if err != nil {
			return
		}
	}
	result, ok := r.data[key]
	if !ok {
		err = ErrKeyNotExistInSession
		return
	}
	return
}
func (r *RedisSession) Del(key string) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()

	delete(r.data, key)
	r.flag = SessionFlagModify
	return
}
func (r *RedisSession) Save() (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()

	if r.flag != SessionFlagModify {
		return
	}
	data, err := json.Marshal(r.data)
	if err != nil {
		return
	}
	conn := r.pool.Get()
	_, err = conn.Do("SET", r.sessionId, string(data))
	if err != nil {
		return
	}
	return
}
