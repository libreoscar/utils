package task

// Note:
// 1. NOT thread-safe
// 2. The order of task execution is guaranteed to be the same as the add order.
type Queue struct {
	tasks []Task
}

func (d *Queue) Push(t Task) {
	d.tasks = append(d.tasks, t)
}

func (d *Queue) ExecAll() {
	for _, t := range d.tasks {
		t()
	}
	d.tasks = nil
}

func (d *Queue) Clear() {
	d.tasks = nil
}

func (d *Queue) Size() int {
	return len(d.tasks)
}
