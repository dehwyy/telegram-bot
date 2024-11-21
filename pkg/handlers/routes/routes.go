package routes

import (
	"github.com/dehwyy/telegram-bot/pkg/logger"
	"github.com/dehwyy/telegram-bot/pkg/state"
	tg "gopkg.in/telebot.v4"
)

// Routes with description

type Route string
type RouteHandler func(tg.Context, *state.State, *logger.Logger) error

func (r Route) String() string { return string(r) }

const (
	RouteDefaultStart Route = "/start"
)

const (
	RouteAdminStart Route = "/admin"
)

const (
	RouteOperatorStart    Route = "/operator"
	RouteOperatorAddOrder Route = "add_order"
)

const (
	RouteProviderStart Route = "/provider"
)

type RouteWithDescription struct {
	Route       Route
	Description string
}

var RoutesOverview []RouteWithDescription = []RouteWithDescription{
	{Route: RouteDefaultStart, Description: "Открыть меню"},
	{Route: RouteAdminStart, Description: "Открыть панель администратора"},
	{Route: RouteOperatorStart, Description: "Открыть панель оператора"},
	{Route: RouteProviderStart, Description: "Открыть панель провайдера"},
}
