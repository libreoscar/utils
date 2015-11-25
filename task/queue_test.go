package task

import (
	"github.com/facebookgo/ensure"
	"testing"
)

func TestQueue(t *testing.T) {
	n := 0

	taskQueue := Queue{}
	taskQueue.Add(func() { n += 1 })
	taskQueue.Add(func() { n += 1 })
	ensure.DeepEqual(t, taskQueue.Size(), 2)

	taskQueue.ExecAll()
	ensure.DeepEqual(t, taskQueue.Size(), 0)
	ensure.DeepEqual(t, n, 2)

	taskQueue.Add(func() { n += 3 })
	taskQueue.ExecAll()
	ensure.DeepEqual(t, n, 5)

	taskQueue.Add(func() { n = 0 })
	taskQueue.Clear()
	ensure.DeepEqual(t, n, 5)
}
