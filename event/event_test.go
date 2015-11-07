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
	// e.EmitAsync()
	// e.EmitAsync()
	e.Emit()
	// e.EmitAsync()
	// e.Wait()
	// e.EmitAsync()
	// e.Wait()
	ensure.DeepEqual(t, counter, 2)
}
