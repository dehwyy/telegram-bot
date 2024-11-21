package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dehwyy/telegram-bot/pkg/db"
	"github.com/dehwyy/telegram-bot/pkg/handlers"
	"github.com/dehwyy/telegram-bot/pkg/handlers/routes"
	"github.com/dehwyy/telegram-bot/pkg/logger"
	"github.com/dehwyy/telegram-bot/pkg/runtime"
	"github.com/dehwyy/telegram-bot/pkg/state"
	tg "gopkg.in/telebot.v4"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	env := runtime.NewEnv()
	l := logger.New(nil)

	// Bot settings.
	pref := tg.Settings{
		Token:  env.Token,
		Poller: &tg.LongPoller{Timeout: 10 * time.Second},
	}

	// Create bot instance.
	b, err := tg.NewBot(pref)
	if err != nil {
		l.Fatal().Err(err).Msg("Failed to create bot")
	}
	defer b.Stop()

	// Process and set available bot commands.
	botCommands := make([]tg.Command, len(routes.RoutesOverview))
	for i, rOverview := range routes.RoutesOverview {
		botCommands[i] = tg.Command{
			Text:        rOverview.Route.String(),
			Description: rOverview.Description,
		}
	}
	b.SetCommands(botCommands)

	// Initialize state & db.
	db := db.NewDatabase()
	state := state.NewState(db)

	// Register handlers.
	handlers.NewDefaultHandler(b.Group(), state, &l)
	handlers.NewOperatorHandler(b.Group(), state, &l)
	// handlers.NewAdminHandler(b.Group(), state, l)

	l.Info().Msg("Starting bot...")
	go b.Start()

	<-ctx.Done()
	l.Info().Msg("Shutting down...")
}
