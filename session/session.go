package session

import (
	"errors"
	"net/http"
	"sync"
)

// Session keeps track of the map between session ids and interface

// TODO should I make Session unexported, only to be created with the
// New() function?
type Session struct {
	mut *sync.Mutex
	// waiting for generics
	session map[string]interface{}
	// sets the cookie name
	// set to "sessionid" in New function
	CookieName string
}

func New() *Session {
	return &Session{
		mut:        &sync.Mutex{},
		session:    make(map[string]interface{}),
		CookieName: "sessionid",
	}
}

// Uses value, ok idiom as return values
func (s *Session) Get(id string) (interface{}, bool) {
	// TODO do I need to lock for reads?
	s.mut.Lock()
	val, ok := s.session[id]
	s.mut.Unlock()
	return val, ok
}

// Uses value, ok idiom as return values
// looks for cookie in CookieName
// returns 'false' if the cookie has not been set or if the value does not exist
func (s *Session) CookieGet(r *http.Request) (interface{}, bool) {
	c, err := r.Cookie(s.CookieName)
	if err != nil {
		return nil, false
	}
	v, ok := s.Get(c.Value)
	return v, ok
}

// grabs the cookie from the request to set the session
// looks for cookie in CookieName
func (s *Session) CookieSet(r *http.Request, val interface{}) error {
	c, err := r.Cookie(s.CookieName)
	if err != nil {
		return errors.New("No cookie set")
	}
	s.Set(c.Value, val)
	return nil
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
