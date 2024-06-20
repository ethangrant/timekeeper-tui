package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Start() {
	p := tea.NewProgram(NewTimeKeeper(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error on start: ", err.Error())
		os.Exit(1)
	}
}
