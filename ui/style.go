package ui

import "github.com/charmbracelet/lipgloss"

var (
	DocStyle          = lipgloss.NewStyle().Margin(1, 2)
	ErrorMessageStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color("9")).Padding(0, 1).Margin(1, 1)
)
