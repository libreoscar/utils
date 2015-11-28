package event

import (
	"reflect"
	"sync"
)

// Best Practice:
// 1) Declare the event as a public member, e.g. FooEvent, in a class;
// 2) Emit the event in the class by self.FooEvent.Emit(...)
// 3) Subscribe the event outside the class by obj.FooEvent.Subscribe(callback)
//    the callback will be executed instantly when FooEvent emits; if you need the callback
//    to be asynchronous executed, use the go keyword manually inside the callback.

type Event struct {
	handlers []reflect.Value
	mu       sync.Mutex
}

func NewEvent() *Event {
	return &Event{
		handlers: nil,
		mu:       sync.Mutex{},
	}
}

func (e *Event) Subscribe(h interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.handlers = append(e.handlers, reflect.ValueOf(h))
}

func (e *Event) Emit(params ...interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if len(e.handlers) > 0 {
		var callArgv []reflect.Value
		for _, p := range params {
			callArgv = append(callArgv, reflect.ValueOf(p))
		}
		for _, h := range e.handlers {
			h.Call(callArgv)
		}
	}
}
