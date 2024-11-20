package handlers

import (
	"github.com/dehwyy/telegram-bot/pkg/handlers/routes"
	tg "gopkg.in/telebot.v4"
)

type handledRoute struct {
	route routes.Route
	h     tg.HandlerFunc
}
