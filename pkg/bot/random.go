package bot

import (
	"fmt"
	apis "github.com/klenje/wikipediarandombot/pkg/apis"
	"golang.org/x/text/language"
	tb "gopkg.in/tucnak/telebot.v2"
)

func (h Handler) OnRandom(m *tb.Message) {
	var pages apis.Pages
	howMany := 2
	lang := "en"
	err := apis.GetRandomWikiPages(&pages, lang, howMany)
	apis.ComputeReadingTime(&pages)
	h.sendMessage(m.Chat, &pages, err)
}

func (h Handler) OnRandomLang(m *tb.Message) {
	if m.Payload == "" {
		h.b.Send(m.Chat, "Please choose a valid Wikipedia language code (e.g. fr)")
	} else {
		var pages apis.Pages
		howMany := 2
		_, err := language.Parse(m.Payload)
		if err != nil {
			h.b.Send(m.Chat, "Invalid language: "+m.Payload)
			return
		}
		lang := m.Payload
		err = apis.GetRandomWikiPages(&pages, lang, howMany)
		apis.ComputeReadingTime(&pages)
		h.sendMessage(m.Chat, &pages, err)
	}
}

func (h Handler) sendMessage(chat *tb.Chat, pages *apis.Pages, err error) {
	if err != nil {
		h.b.Send(chat, err)
		return
	}
	for _, page := range pages.Items {
		result := fmt.Sprintf("%s (%d - %d minutes read) <a href=\"%s\">link</a>\n", page.Title, int(page.Length), int(page.TimeToRead), page.Fullurl)
		h.b.Send(chat, result)
	}
}
