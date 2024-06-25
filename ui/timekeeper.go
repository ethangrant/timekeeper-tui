package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethangrant/timekeeper/tasks"
)

var timeKeeperModel timeKeeper

type timeKeeper struct {
	help help.Model
	list list.Model
	err  error
}

func NewTimeKeeper() *timeKeeper {
	timeKeeperModel = timeKeeper{
		help: help.New(),
		list: NewList("Tasks for: ", time.Now()),
	}

	return &timeKeeperModel
}

func (t *timeKeeper) Init() tea.Cmd {
	tea.LogToFile("debug.log", "tk init")
	return nil
}

func (t *timeKeeper) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return t, tea.Quit
		case key.Matches(msg, keys.New):
			return NewForm().Update(msg)
		}

	case tea.WindowSizeMsg:
		h, v := DocStyle.GetFrameSize()
		t.list.SetSize(msg.Width-h, msg.Height-v)
	case TaskFormSubmittedMsg:
		return t, t.addTask(msg.task)
	case TaskFormSubmittedErrorMsg:
		t.err = msg.err
		return t, nil
	}

	var cmd tea.Cmd
	t.list, cmd = t.list.Update(msg)

	return t, cmd
}

func (t *timeKeeper) View() string {
	var errorMsg string

	if t.err != nil {
		errorMsg = t.err.Error()
	}

	// reset the err
	t.err = nil

	return lipgloss.JoinVertical(
		lipgloss.Left,
		errorMsg,
		DocStyle.Render(t.list.View()),
		t.help.View(keys),
	)
}

func (t *timeKeeper) addTask(task tasks.Task) tea.Cmd {
	return t.list.InsertItem(0, task)
}
