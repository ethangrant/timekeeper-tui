package tasks

import (
	"time"

	"github.com/ethangrant/timekeeper/stopwatch"
)

type Task struct {
	title string
	desc  string
	Timer *stopwatch.Model
	// created time.Time
}

func New(title string, desc string, duration time.Duration) Task {
	return Task{
		title: title,
		desc:  desc,
		Timer: stopwatch.New(0),
	}
}

func (t Task) Title() string       { return t.title }
func (t Task) Description() string { return t.desc }
func (t Task) FilterValue() string { return t.title }
