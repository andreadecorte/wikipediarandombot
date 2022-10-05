package bot

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func (h Handler) OnHelp(m *tb.Message) {
	HelpText := `
	<i>Welcome!</i>
	How to use:
	<b>/random</b>
	<b>/randomlang langcode</b>
	`
	h.b.Send(m.Chat, HelpText)
}
