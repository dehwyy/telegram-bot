package handlers

import (
	"fmt"

	"github.com/dehwyy/telegram-bot/pkg/handlers/routes"
	"github.com/dehwyy/telegram-bot/pkg/logger"
	"github.com/dehwyy/telegram-bot/pkg/middleware"
	"github.com/dehwyy/telegram-bot/pkg/state"
	tg "gopkg.in/telebot.v4"
)

var (
	addOrderButton = tg.InlineButton{
		Text:   "Добавить заявку",
		Unique: routes.RouteOperatorAddOrder.String(),
	}

	operatorMenu = &tg.ReplyMarkup{
		InlineKeyboard: [][]tg.InlineButton{{addOrderButton}, {{
			Text: "Заявки",
			Data: "orders",
		}}},
	}

	selectBankMenu = &tg.ReplyMarkup{
		InlineKeyboard: [][]tg.InlineButton{{{
			Text: "Сбербанк",
			Data: "sberbank",
		}, {
			Text: "Тинькофф",
			Data: "tinkoff",
		}, {
			Text: "Сбербанк СБП",
			Data: "sberbank_sbp",
		}}},
	}
)

func NewOperatorHandler(group *tg.Group, state *state.State, l *logger.Logger) {
	g := routes.NewGroup(group)

	// TODO: impl
	g.Group.Use(middleware.RoleMiddleware)

	g.Add(routes.RouteOperatorStart, operatorMenuHandler)
	g.AddCallback(&addOrderButton, operatorAddOrderHandler)

	g.Handle(state, l)
}

func operatorAddOrderHandler(ctx tg.Context, state *state.State, l *logger.Logger) error {
	return ctx.Edit("Выберите банк", selectBankMenu)
}

func operatorMenuHandler(ctx tg.Context, state *state.State, l *logger.Logger) error {

	completed_orders := 3
	active_orders := 2
	cancelled_orders := 10

	user_balance := 30624
	user_commision := 5

	return ctx.Send(
		fmt.Sprintf(
			"<b>Меню оператора</b>\n\n"+
				"<b>Кол-во заявок в работе:</b> %d\n"+
				"<b>Кол-во обработанных заявок:</b> %d\n"+
				"<b>Кол-во активных заявок:</b> %d\n"+
				"<b>Кол-во отмененных заявок:</b> %d\n"+
				"<b>Баланс пользователя:</b> %d\n"+
				"<b>Комиссия:</b> %d%%\n",
			active_orders,
			completed_orders,
			active_orders,
			cancelled_orders,
			user_balance,
			user_commision,
		),
		operatorMenu, tg.ModeHTML,
	)
}
