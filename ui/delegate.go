package ui

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ethangrant/timekeeper/stopwatch"
	"github.com/ethangrant/timekeeper/taskdb"
	"github.com/ethangrant/timekeeper/tasks"
)

type UpdateTaskDurationMsg struct {
	task tasks.Task
}

type UpdateTaskDurationErrMsg struct {
	err error
}

// Implementing a custom item delegate so we can add the timer.
type itemDelegate struct{}

func (d itemDelegate) Height() int  { return list.NewDefaultDelegate().Height() }
func (d itemDelegate) Spacing() int { return list.NewDefaultDelegate().Spacing() }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	task, _ := m.SelectedItem().(tasks.Task)
	switch msg := msg.(type) {
	case stopwatch.StartStopMsg:
		return updateTaskDuration(task)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.StartTimer, keys.StopTimer):
			keys.StartTimer.SetEnabled(!task.Timer.Running())
			keys.StopTimer.SetEnabled(task.Timer.Running())
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

// TODO: Handle this message somewhere probably timekeeper
func updateTaskDuration(task tasks.Task) tea.Cmd {
	return func () tea.Msg {
		ctx := context.Background()
		queries := taskdb.New(DbConn)
		_, err := queries.UpdateTaskDuration(
			ctx, taskdb.UpdateTaskDurationParams{
				Duration: int64(task.Timer.Elapsed()),
				ID: task.Id(),
			},
		)

		if err != nil {
			return UpdateTaskDurationErrMsg{err: err}
		}

		return UpdateTaskDurationMsg{task: task}
	}
}
