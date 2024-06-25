package ui

import (
	"database/sql"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"

	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

// setup a global db instance
var DbConn *sql.DB
func dbInit() *sql.DB {
	filename := "/tmp/timekeeper.db"
	_, err := os.Stat(filename)
	if err != nil {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("failed to create timekeeper.db")
			os.Exit(1)
		}

		defer file.Close()
	}

	conn, err := sql.Open("sqlite3", filename)
	if err != nil {
		fmt.Println("database connection failed: ", err.Error())
		os.Exit(1)
	}

	DbConn = conn

	driver, err := sqlite3.WithInstance(DbConn, &sqlite3.Config{})
	if err != nil {
		fmt.Println("error getting sqlite drive with instance: ", err.Error())
		os.Exit(1)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"sqlite3",
		driver,
	)

	if err != nil {
		fmt.Println("error migrate new with db instance: ", err.Error())
		os.Exit(1)
	}

	m.Up()

	return DbConn
}

func Start() {
	dbInit()
	p := tea.NewProgram(NewTimeKeeper(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error on start: ", err.Error())
		os.Exit(1)
	}
}
