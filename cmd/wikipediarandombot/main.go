package main

import (
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"

	"wikipediarandombot/pkg/bot"

	log "github.com/sirupsen/logrus"
)

func main() {

	var (
		verbose              = true
		useWebhook           = os.Getenv("WEBHOOK")
		port                 = os.Getenv("PORT")
		publicURL            = os.Getenv("PUBLIC_URL")
		tgToken              = os.Getenv("TELEGRAM_TOKEN")
		poller     tb.Poller = nil
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	poller = &tb.LongPoller{Timeout: 10 * time.Second}
	// Webhook needs a public URL
	if useWebhook != "" && publicURL != "" {
		log.Info("Using webhook")
		poller = webhook
	}

	if tgToken == "" {
		log.Fatal("Telegram token missing, check README")
	}

	b, err := tb.NewBot(tb.Settings{
		Token:     tgToken,
		Poller:    poller,
		Verbose:   verbose,
		ParseMode: tb.ModeHTML,
	})

	if err != nil {
		log.Fatal(err)
	}

	if useWebhook == "" {
		// If we switch from webhook to poller
		// we may need to delete the webhook
		b.RemoveWebhook()
	}

	h := bot.New(bot.Config{
		Bot: b,
	})

	b.Handle("/help", h.OnHelp)
	b.Handle("/start", h.OnHelp)
	b.Handle("/random", h.OnRandom)
	b.Handle("/randomlang", h.OnRandomLang)

	b.Handle(tb.OnText, h.OnHelp)
	//b.Handle(tb.OnQuery, h.OnInlineQuery)

	log.Info("Start bot")
	b.Start()
}
