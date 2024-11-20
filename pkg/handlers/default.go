package handlers

import (
	"github.com/dehwyy/telegram-bot/pkg/handlers/routes"
	tg "gopkg.in/telebot.v4"
)

func NewDefaultHandler(group *tg.Group) {
	handledRoutes := []handledRoute{
		{routes.RouteDefaultStart, defaultHandleStart},
	}

	for _, r := range handledRoutes {
		group.Handle(r.route.String(), r.h)
	}
}

func defaultHandleStart(ctx tg.Context) error {
	return ctx.Send("Default handler message")
}
