package bot

import (
	"fmt"

	"golang.org/x/text/language"
	tele "gopkg.in/telebot.v3"

	"wikipediarandombot/pkg/apis"
)

func OnRandom(c tele.Context) (err error) {
	var pages apis.Pages
	howMany := 2
	lang := "en"
	err = apis.GetRandomWikiPages(&pages, lang, howMany)
	if err != nil {
		return
	}
	err = apis.ComputeReadingTime(&pages)
	if err != nil {
		return
	}
	sendMessage(c, &pages)
	return
}

func OnRandomLang(c tele.Context) (err error) {
	if c.Message().Payload == "" {
		c.Send("Please choose a valid Wikipedia language code (e.g. fr)")
		return
	}

	var pages apis.Pages
	howMany := 2
	_, err = language.Parse(c.Message().Payload)
	if err != nil {
		c.Send("Invalid language: " + c.Message().Payload)
		return
	}
	lang := c.Message().Payload
	err = apis.GetRandomWikiPages(&pages, lang, howMany)
	if err != nil {
		return
	}
	err = apis.ComputeReadingTime(&pages)
	if err != nil {
		return
	}
	sendMessage(c, &pages)
	return
}

func sendMessage(c tele.Context, pages *apis.Pages) {
	for _, page := range pages.Items {
		result := fmt.Sprintf("%s (%d - %d minutes read) <a href=\"%s\">link</a>\n", page.Title, int(page.Length), page.TimeToRead, page.Fullurl)
		c.Send(result)
	}
}
