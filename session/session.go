package session

import (
	"sync"
)

// Session keeps track of the map between session ids and interface

type Session struct {
	mut *sync.Mutex
	// waiting for generics
	session map[string]interface{}
}

func New() *Session {
	return &Session{
		mut: &sync.Mutex{},
		session: make(map[string]interface{}),
	}
}

// Uses value, ok idiom as return values
func (s *Session) Get(id string) (val interface{}, ok bool) {
	// TODO do I need to lock for reads?
	s.mut.Lock()
	val, ok = s.session[id]
	s.mut.Unlock()
	return
}

func (s *Session) Set(id string, val interface{}) {
	s.mut.Lock()
	s.session[id] = val
	s.mut.Unlock()
}

func (s *Session) Unset(id string) {
	s.mut.Lock()
	delete(s.session, id)
	s.mut.Unlock()
}
