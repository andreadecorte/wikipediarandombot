package bot

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

type Handler struct {
	b *tb.Bot
}

type Config struct {
	Bot *tb.Bot
}

func New(c Config) Handler {
	h := Handler{
		b: c.Bot,
	}

	return h
}
