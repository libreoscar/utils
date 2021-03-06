package event

import (
	"reflect"
	"sync"
)

// Best Practices:
// 1. Declare the event as a public member, e.g. FooEvent, in a class;
// 2. Emit the event in the class by self.FooEvent.Emit(...)
// 3. Subscribe the event outside the class by obj.FooEvent.Subscribe(callback)
//    the callback will be executed instantly when FooEvent emits; if you need the callback
//    to be asynchronous executed, use the go keyword manually inside the callback.
// 4. When the callback of a subscriber returns UNSUBSCRIBE, this subscriber will be removed from
//    the subscriber list of this event, and thus will not receive more event.

// TODO:
// 1. Optimize this implementation. We can mark the handlers to be deleted as nil in the
//    first pass, and shift non-nil handlers the front in the second pass, and then shrink
//    the slice to fit.
// 2. Store the last event (can also be initially in the beginning). Allow new handler to
//    be called with the last event when subscribing.

type Event struct {
	handlers []reflect.Value
	mu       sync.Mutex
}

const (
	UNSUBSCRIBE = true
)

func NewEvent() *Event {
	return &Event{
		handlers: nil,
		mu:       sync.Mutex{},
	}
}

func (e *Event) Size() int {
	e.mu.Lock()
	defer e.mu.Unlock()
	return len(e.handlers)
}

func (e *Event) Subscribe(h interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.handlers = append(e.handlers, reflect.ValueOf(h))
}

func (e *Event) Emit(params ...interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()

	toRemove := make(map[int]bool)
	if len(e.handlers) > 0 {
		var callArgv []reflect.Value
		for _, p := range params {
			callArgv = append(callArgv, reflect.ValueOf(p))
		}
		for idx, h := range e.handlers {
			rst := h.Call(callArgv)
			// if rst is true, the handler will be removed
			if len(rst) == 1 && rst[0].Interface() == interface{}(UNSUBSCRIBE) {
				toRemove[idx] = true
			}
		}
	}

	if len(toRemove) > 0 {
		var afterRemove []reflect.Value = nil
		for idx, h := range e.handlers {
			if !toRemove[idx] {
				afterRemove = append(afterRemove, h)
			}
		}
		e.handlers = afterRemove
	}
}
