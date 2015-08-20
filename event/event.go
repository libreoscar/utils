package event

import (
	"reflect"
	"sync"
)

type Event struct {
	sync.Mutex
	handlers []reflect.Value
}

func (s *Event) Subscribe(h interface{}) {
	s.Lock()
	defer s.Unlock()

	s.handlers = append(s.handlers, reflect.ValueOf(h))
}

func (s *Event) Emit(params ...interface{}) {
	s.Lock()
	defer s.Unlock()

	if len(s.handlers) > 0 {
		var callArgv []reflect.Value
		for _, p := range params {
			callArgv = append(callArgv, reflect.ValueOf(p))
		}
		for _, h := range s.handlers {
			h.Call(callArgv)
		}
	}
}
