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
