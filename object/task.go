package object

import "time"

type Task struct {
	Title   string    `json:title`
	Content string    `json:content`
	DueDate time.Time `json:DueDate`
}

func NewTask(t string, c string, d time.Time) *Task {
	return &Task{
		Title:   t,
		Content: c,
		DueDate: d,
	}
}

func (t *Task) Init() {
	t.Title = "TaskのSampleTitle"
	t.Content = "TaskのSampleContent"
	t.DueDate = time.Now()
}
