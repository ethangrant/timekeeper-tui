package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ethangrant/timekeeper/ui"
)

func main() {
	tea.LogToFile("debug.log", "tk")
	ui.Start()
}
