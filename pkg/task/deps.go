package task

import ()

func (t *Task) AddDep(task *Task) {
	t.resolvedDeps = append(t.resolvedDeps, task)
}
