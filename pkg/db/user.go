package db

import "github.com/dehwyy/telegram-bot/pkg"

type UserDb struct{}

type User struct {
	Id       int64
	Name     string
	UserRole pkg.UserRole
}

func (userDb *UserDb) GetUserById(id int64) (*User, error) {
	return &User{Id: id, Name: "dehwyy", UserRole: pkg.Operator}, nil
}
