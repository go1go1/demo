package session

import (
	"sync"
)

type MemorySession struct {
	sessionId string
	data      map[string]interface{}
	rwLock    sync.RWMutex
}

func NewMemorySession(sessionId string) *MemorySession {
	s := &MemorySession{
		sessionId: sessionId,
		data:      make(map[string]interface{}, 8),
	}
	return s
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()

	m.data[key] = value
	return
}

func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	value, ok := m.data[key]
	if !ok {
		err = ErrKeyNotExistInSession
		return
	}
	return
}

func (m *MemorySession) Del(key string) (err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()

	delete(m.data, key)
	return
}

func (m *MemorySession) Save() (err error) {
	return
}
