package bot

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func (h Handler) OnHelp(m *tb.Message) {
	HELP_TEXT := `
	<i>Welcome!</i>
	`
	h.b.Send(m.Chat, HELP_TEXT)
}