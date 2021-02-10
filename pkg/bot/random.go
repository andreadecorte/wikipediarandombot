package bot

import (
	apis "github.com/klenje/wikipediarandombot/pkg/apis"
	tb "gopkg.in/tucnak/telebot.v2"
	"fmt"
)

func (h Handler) OnRandom(m *tb.Message) {
	var pages apis.Pages
	howMany := 2
	lang := "en"
	err := apis.GetWiki(&pages, lang, howMany)
	if err != nil {
		h.b.Send(m.Chat, err)
	}
	for _, page := range pages.Items {
		result := fmt.Sprintf("%s (%f) <a href=\"%s\">link</a>\n", page.Title, page.Length, page.Fullurl)
		h.b.Send(m.Chat, result)
	}
}