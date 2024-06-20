package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type form struct {
	help  help.Model
	title textinput.Model
	desc  textarea.Model
	timer stopwatch.Model
}

func NewForm() form {
	form := form{
		help:  help.New(),
		title: textinput.New(),
		desc:  textarea.New(),
		timer: stopwatch.New(),
	}

	form.title.Placeholder = "Task title"
	form.desc.Placeholder = "What have you been working on?"
	form.title.Focus()

	return form
}

func (f form) Init() tea.Cmd {
	return nil
}

func (f form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.StartTimer, keys.StopTimer):
			keys.StartTimer.SetEnabled(!f.timer.Running())
			keys.StopTimer.SetEnabled(f.timer.Running())
			return f, f.timer.Toggle()
		case key.Matches(msg, keys.ResetTimer):
			f.timer.Reset()
		case key.Matches(msg, keys.Quit):
			return f, tea.Quit

		case key.Matches(msg, keys.Back):
			return timeKeeperModel.Update(nil)
		case key.Matches(msg, keys.Enter):
			if f.title.Focused() {
				f.title.Blur()
				f.desc.Focus()
				return f, textarea.Blink
			}
			// Return the completed form as a message.
			return timeKeeperModel.Update(f)
		}
	}
	var cmd tea.Cmd
	if f.title.Focused() {
		f.title, cmd = f.title.Update(msg)
	} else {
		f.desc, cmd = f.desc.Update(msg)
	}

	// Update the stopwatch
	var swCmd tea.Cmd
	f.timer, swCmd = f.timer.Update(msg)

	return f, tea.Batch(cmd, swCmd)
}

func (f form) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		"Create a new task",
		f.title.View(),
		f.desc.View(),
		f.timer.View(),
		f.help.View(keys))
}