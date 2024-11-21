package middleware

import (
	tg "gopkg.in/telebot.v4"
)

// TODO:
func RoleMiddleware(next tg.HandlerFunc) tg.HandlerFunc {
	return func(ctx tg.Context) error {
		return next(ctx)
	}
}
