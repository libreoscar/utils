package event

import (
	"github.com/facebookgo/ensure"
	"testing"
	"time"
)

func TestEvent(t *testing.T) {
	counter := 0

	e := NewEvent()
	e.Subscribe(func() {
		counter = counter + 1
		time.Sleep(10 * time.Millisecond)
	})
	e.Subscribe(func() {
		counter = counter + 2
	})

	e.Emit()
	e.EmitAsync()
	e.EmitAsync()
	e.Emit()
	e.EmitAsync()
	e.Wait()
	e.EmitAsync()
	e.Wait()
	ensure.DeepEqual(t, counter, 15)
}
