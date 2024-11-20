package handlers

import (
	"github.com/dehwyy/telegram-bot/pkg/handlers/routes"
	tg "gopkg.in/telebot.v4"
)

func NewAdminHandler(group *tg.Group) {
	handledRoutes := []handledRoute{
		{routes.RouteAdminStart, adminHandleStart},
	}

	for _, r := range handledRoutes {
		group.Handle(r.route.String(), r.h)
	}
}

func adminHandleStart(ctx tg.Context) error {
	return ctx.Send("admin message")
}
