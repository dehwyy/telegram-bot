package state

import "github.com/dehwyy/telegram-bot/pkg/db"

type State struct {
	DB *db.Database
}

func NewState(db *db.Database) *State {
	return &State{DB: db}
}
