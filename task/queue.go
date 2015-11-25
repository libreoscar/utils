package task

type Queue struct {
	tasks []Task
}

func (d *Queue) Add(t Task) {
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
