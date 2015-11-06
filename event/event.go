package event

import (
	"reflect"
)

type Event struct {
	handlers     []reflect.Value
	blockingChan chan struct{}
}

func NewEvent() *Event {
	return &Event{
		blockingChan: make(chan struct{}, 1),
	}
}

func (e *Event) Subscribe(h interface{}) {
	e.blockingChan <- struct{}{}
	defer func() {
		<-e.blockingChan
	}()

	e.handlers = append(e.handlers, reflect.ValueOf(h))
}

func (e *Event) Emit(params ...interface{}) {
	e.blockingChan <- struct{}{}
	defer func() {
		<-e.blockingChan
	}()

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

func (e *Event) EmitAsync(params ...interface{}) {
	e.blockingChan <- struct{}{}
	go func() {
		defer func() {
			<-e.blockingChan
		}()
		if len(e.handlers) > 0 {
			var callArgv []reflect.Value
			for _, p := range params {
				callArgv = append(callArgv, reflect.ValueOf(p))
			}
			for _, h := range e.handlers {
				h.Call(callArgv)
			}
		}
	}()
}

// wait for all handlers to finish
func (e *Event) Wait() {
	e.blockingChan <- struct{}{}
	defer func() {
		<-e.blockingChan
	}()
}
