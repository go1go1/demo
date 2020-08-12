package tools

import (
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	m map[interface{}]interface{}
}

func (m *SafeMap) init() {
	if m.m == nil {
		m.m = make(map[interface{}]interface{})
	}
}

func (m *SafeMap) UnsafeGet(key interface{}) interface{} {
	if m.m == nil {
		return nil
	} else {
		return m.m[key]
	}
}

func (m *SafeMap) Get(key interface{}) interface{} {
	m.RLock()
	defer m.RUnlock()
	return m.UnsafeGet(key)
}

func (m *SafeMap) UnsafeSet(key interface{}, value interface{}) {
	m.init()
	m.m[key] = value
}

func (m *SafeMap) Set(key interface{}, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.UnsafeSet(key, value)
}

func (m *SafeMap) TestAndSet(key interface{}, value interface{}) interface{} {
	m.Lock()
	defer m.Unlock()

	m.init()

	if v, ok := m.m[key]; ok {
		return v
	} else {
		m.m[key] = value
		return nil
	}
}

func (m *SafeMap) UnsafeDel(key interface{}) {
	m.init()
	delete(m.m, key)
}

func (m *SafeMap) Del(key interface{}) {
	m.Lock()
	defer m.Unlock()
	m.UnsafeDel(key)
}

func (m *SafeMap) UnsafeLen() int {
	if m.m == nil {
		return 0
	} else {
		return len(m.m)
	}
}

func (m *SafeMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return m.UnsafeLen()
}

func (m *SafeMap) UnsafeRange(f func(interface{}, interface{})) {
	if m.m == nil {
		return
	}
	for k, v := range m.m {
		f(k, v)
	}
}

func (m *SafeMap) RLockRange(f func(interface{}, interface{})) {
	m.RLock()
	defer m.RUnlock()
	m.UnsafeRange(f)
}

func (m *SafeMap) LockRange(f func(interface{}, interface{})) {
	m.Lock()
	defer m.Unlock()
	m.UnsafeRange(f)
}
