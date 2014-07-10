package session

import (
	"sync"
)

// Session keeps track of the map between session ids and bools

type Session struct {
	mut *sync.Mutux
	session map[string]bool
}

func New() *Session {
	return &Session{
		mut: &sync.Mutex{},
		session: make(map[string]bool),
	}
}

// Uses value, ok idiom as return values
func (s *Session) Get(id string) (val bool, ok bool) {
	// TODO do I need to lock for reads?
	s.mut.Lock()
	val, ok = s.session[id]
	s.mut.Unlock()
	return
}

func (s *Session) Set(id string, val bool) {
	s.mut.Lock()
	s.session[id] = val
	s.mut.Unlock()
}
