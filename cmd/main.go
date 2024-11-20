package main

import (
	"log"
	"time"

	"github.com/dehwyy/telegram-bot/pkg/runtime"
	tg "gopkg.in/telebot.v4"
)

func main() {
	env := runtime.NewEnv()
	pref := tg.Settings{
		Token:  env.Token,
		Poller: &tg.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tg.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	b.Handle("/hello", func(ctx tg.Context) error {
		p := ctx.Message().Text
		return ctx.Send(p)
	})

	b.Handle("/start", func(ctx tg.Context) error {
		return ctx.Send("иди в жопу")
	})

	b.Start()
}
