package app

import (
	"gopkg.in/mgo.v2"
)

// App struct implemented domain layer.

type App struct {
	// Connect with percona database.
	DB *mgo.Database
}

func New(db *mgo.Database) *App {
	r := &App{
		DB: db,
	}
	return r
}
