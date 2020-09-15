package session

import (
	"sync"

	"github.com/google/uuid"
)

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwLock     sync.RWMutex
}

func NewMemorySessionMgr() SessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return sr
}

func (s *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

func (s *MemorySessionMgr) GetSession(sessionId string) (session Session, err error) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()

	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = ErrSessionNotExist
		return
	}
	return
}

func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	id := uuid.New()
	sessionId := id.String()

	session = NewMemorySession(sessionId)
	s.sessionMap[sessionId] = session
	return
}
