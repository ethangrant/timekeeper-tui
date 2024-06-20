package tasks

// import "time"

type Task struct {
	title string
	desc  string
	// totalTime string
	// created time.Time
}

func New(title string, desc string) Task {
	return Task{
		title: title,
		desc:  desc,
		// totalTime: total,
		// created: created,
	}
}

func (t Task) Title() string       { return t.title }
func (t Task) Description() string { return t.desc }
func (t Task) FilterValue() string { return t.title }
