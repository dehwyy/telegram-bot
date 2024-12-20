package handlers

import (
	"fmt"

	"github.com/dehwyy/telegram-bot/pkg"
	"github.com/dehwyy/telegram-bot/pkg/handlers/routes"
	"github.com/dehwyy/telegram-bot/pkg/logger"
	"github.com/dehwyy/telegram-bot/pkg/state"
	tg "gopkg.in/telebot.v4"
)

func NewDefaultHandler(group *tg.Group, state *state.State, l *logger.Logger) {
	g := routes.NewGroup(group)
	g.Add(routes.RouteDefaultStart, defaultStartHandler)

	g.Handle(state, l)
}

func defaultStartHandler(ctx tg.Context, state *state.State, l *logger.Logger) error {
	userId := ctx.Sender().ID
	l.Info().Msgf("UserId: %d", userId)

	user, _ := state.DB.User().GetUserById(userId)
	l.Info().Msgf("User: %s", user.Name)

	switch user.UserRole {
	case pkg.Operator:
		return operatorMenuHandler(ctx, state, l)
	}

	token := ctx.Message().Payload
	l.Info().Msgf("Token: %s", token)

	return ctx.Reply(fmt.Sprintf("Hello, %s!", ctx.Sender().Username), tg.ModeMarkdown)
}
