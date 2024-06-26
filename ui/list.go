package ui

import (
	"context"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/ethangrant/timekeeper/taskdb"
	"github.com/ethangrant/timekeeper/tasks"
)

// return a new bubble list using tasks in the database
func NewList(title string, date time.Time) list.Model {
	bubbleList := []list.Item{}
	ctx := context.Background()
	queries := taskdb.New(DbConn)
	tsks, err := queries.GetAllTasksByDate(ctx, date.Format("2006-01-02"))
	if err != nil {
		fmt.Println("failed to load tasks")
	}

	for _, tsk := range tsks {
		bubbleList = append(bubbleList, tasks.New(tsk.ID, tsk.Title, tsk.Desc, time.Duration(tsk.Duration)))
	}

	lst := list.New(bubbleList, NewItemDelegate(), 0, 0)
	lst.Title = title + date.Format("02/01/2006")

	return lst
}
