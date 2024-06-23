package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ethangrant/timekeeper/tasks"
)

var timeKeeperModel timeKeeper

type timeKeeper struct {
	list list.Model
}

func NewTimeKeeper() *timeKeeper {
	now := time.Now()
	formatted := now.Format("02/01/2006")

	timeKeeperModel = timeKeeper{
		NewList("Tasks for: " + formatted),
	}

	//TODO: some logic to grab tasks from txt file.

	timeKeeperModel.list.InsertItem(0, tasks.New("CUT-1337", "SOME DESC"))

	return &timeKeeperModel
}

func (t *timeKeeper) Init() tea.Cmd {
	tea.LogToFile("debug.log", "tk init")
	return nil
}

func (t *timeKeeper) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return t, tea.Quit
		case "n":
			return NewForm().Update(msg)
		}
	case tea.WindowSizeMsg:
		h, v := DocStyle.GetFrameSize()
		t.list.SetSize(msg.Width-h, msg.Height-v)
	}
	var cmd tea.Cmd
	t.list, cmd = t.list.Update(msg)

	return t, cmd
}

func (t *timeKeeper) View() string {
	return DocStyle.Render(t.list.View())
}
