package bot

import (
	tele "gopkg.in/telebot.v3"
)

func OnHelp(c tele.Context) error {
	HelpText := `
	<i>Welcome!</i>
	How to use:
	<b>/random</b>
	<b>/randomlang langcode</b>
	`
	return c.Send(HelpText)
}
