package ui

import "github.com/charmbracelet/bubbles/list"

func NewList(title string) list.Model {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = title

	return l
}
