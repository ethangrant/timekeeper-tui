package tasks

import (
	"time"

	"github.com/ethangrant/timekeeper/stopwatch"
)

type Task struct {
	id    int64
	title string
	desc  string
	Timer *stopwatch.Model
	// created time.Time
}

func New(id int64, title string, desc string, duration time.Duration) Task {
	return Task{
		id:    id,
		title: title,
		desc:  desc,
		Timer: stopwatch.New(0),
	}
}

func (t Task) Id() int64           { return t.id }
func (t Task) Title() string       { return t.title }
func (t Task) Description() string { return t.desc }
func (t Task) FilterValue() string { return t.title }
