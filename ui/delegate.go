package ui

import (
	"fmt"
	"io"
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ethangrant/timekeeper/tasks"
)

// Implementing a custom item delegate so we can add the timer.
type itemDelegate struct{}

func (d itemDelegate) Height() int  { return list.NewDefaultDelegate().Height() }
func (d itemDelegate) Spacing() int { return list.NewDefaultDelegate().Spacing() }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	task, _ := m.SelectedItem().(tasks.Task)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.StartTimer, keys.StopTimer):
			keys.StartTimer.SetEnabled(!task.Timer.Running())
			keys.StopTimer.SetEnabled(task.Timer.Running())
			log.Default().Println("Start timer")
			return task.Timer.Toggle()
		case key.Matches(msg, keys.ResetTimer):
			return task.Timer.Reset()
		}
	}

	var swCmd tea.Cmd
	log.Default().Print(msg)
	task.Timer, swCmd = task.Timer.Update(msg)

	return swCmd
}

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	list.NewDefaultDelegate().Render(w, m, index, listItem)
	if task, ok := listItem.(tasks.Task); ok {
		fmt.Fprintf(w, "%s%s", "\n", task.Timer.View())
	}
}

func NewItemDelegate() *itemDelegate {
	return &itemDelegate{}
}
