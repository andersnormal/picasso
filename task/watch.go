package task

func (t *Task) ShouldWatch() bool {
	return len(t.Watch.Paths) > 0 || len(t.Watch.Ignore) > 0
}
