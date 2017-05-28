package event

import (
	"github.com/facebookgo/ensure"
	"testing"
)

func TestEvent(t *testing.T) {
	counter := 0

	e := NewEvent()
	e.Subscribe(func() {
		counter = counter + 1
	})

	e.Emit()
	e.Emit()
	ensure.DeepEqual(t, counter, 2)
}

func TestRemove(t *testing.T) {
	counter := 0
	e := NewEvent()
	e.Subscribe(func() {
		counter += 2
	})
	e.Subscribe(func() bool {
		counter += 1
		return UNSUBSCRIBE
	})

	e.Emit()
	e.Emit()
	ensure.DeepEqual(t, counter, 5)
}
