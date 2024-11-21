package routes

import (
	"github.com/dehwyy/telegram-bot/pkg/logger"
	"github.com/dehwyy/telegram-bot/pkg/state"
	tg "gopkg.in/telebot.v4"
)

type routeWithHandler struct {
	Route   Route
	Handler RouteHandler
}

type callbackHandler struct {
	Endpoint tg.CallbackEndpoint
	Handler  RouteHandler
}

type Group struct {
	Group     *tg.Group
	Routes    []routeWithHandler
	Callbacks []callbackHandler
}

func NewGroup(group *tg.Group) *Group {
	return &Group{Group: group}
}

func (g *Group) Add(route Route, h RouteHandler) {
	g.Routes = append(g.Routes, routeWithHandler{Route: route, Handler: h})
}

func (g *Group) AddCallback(endpoint tg.CallbackEndpoint, h RouteHandler) {
	g.Callbacks = append(g.Callbacks, callbackHandler{Endpoint: endpoint, Handler: h})
}

func (g *Group) Handle(state *state.State, l *logger.Logger) {
	for _, r := range g.Routes {
		g.Group.Handle(r.Route.String(), func(ctx tg.Context) error {
			return r.Handler(ctx, state, l)
		})
	}

	for _, r := range g.Callbacks {
		g.Group.Handle(r.Endpoint, func(ctx tg.Context) error {
			return r.Handler(ctx, state, l)
		})
	}
}
