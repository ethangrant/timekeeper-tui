package ui

import (
	"context"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethangrant/timekeeper/taskdb"
	"github.com/ethangrant/timekeeper/tasks"
)

type form struct {
	heading string
	help    help.Model
	title   textinput.Model
	desc    textarea.Model
}

type TaskFormSubmittedMsg struct {
	task tasks.Task
}

type TaskFormSubmittedErrorMsg struct {
	err error
}

func NewForm() *form {
	form := form{
		heading: "Create a new task",
		help:    help.New(),
		title:   textinput.New(),
		desc:    textarea.New(),
	}

	form.title.Placeholder = "Ticket"
	form.desc.Placeholder = "What have you been working on?"
	form.title.Focus()

	return &form
}

func (f *form) Init() tea.Cmd {
	return nil
}

func (f *form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Back):
			return timeKeeperModel.Update(nil)
		case key.Matches(msg, keys.Enter):
			if f.title.Focused() {
				f.title.Blur()
				f.desc.Focus()
				return f, textarea.Blink
			}

			t, cmd := timeKeeperModel.Update(f)
			return t, tea.Batch(cmd, taskFormSubmitted(f.title.Value(), f.desc.Value()))
		}
	}
	var cmd tea.Cmd
	if f.title.Focused() {
		f.title, cmd = f.title.Update(msg)
	} else {
		f.desc, cmd = f.desc.Update(msg)
	}

	return f, tea.Batch(cmd)
}

func (f *form) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		f.heading,
		f.title.View(),
		f.desc.View(),
		f.help.View(keys))
}

// cmd to insert new task
func taskFormSubmitted(title string, desc string) tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()
		queries := taskdb.New(DbConn)

		insertedTask, err := queries.CreateTask(ctx, taskdb.CreateTaskParams{
			Title:    title,
			Desc:     desc,
			Duration: 0,
		})

		if err != nil {
			return TaskFormSubmittedErrorMsg{err: err}
		}

		tsk := tasks.New(insertedTask.ID, insertedTask.Title, insertedTask.Desc, time.Duration(insertedTask.Duration))

		return TaskFormSubmittedMsg{task: tsk}
	}
}
