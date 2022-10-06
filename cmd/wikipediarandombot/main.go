package main

import (
	"os"
	"time"

	tele "gopkg.in/telebot.v3"

	"wikipediarandombot/pkg/bot"

	log "github.com/sirupsen/logrus"
)

func main() {

	var (
		verbose                = true
		useWebhook             = os.Getenv("WEBHOOK")
		port                   = os.Getenv("PORT")
		publicURL              = os.Getenv("PUBLIC_URL")
		tgToken                = os.Getenv("TELEGRAM_TOKEN")
		poller     tele.Poller = nil
	)

	webhook := &tele.Webhook{
		Listen:   ":" + port,
		Endpoint: &tele.WebhookEndpoint{PublicURL: publicURL},
	}

	poller = &tele.LongPoller{Timeout: 10 * time.Second}
	// Webhook needs a public URL
	if useWebhook != "" && publicURL != "" {
		log.Info("Using webhook")
		poller = webhook
	}

	if tgToken == "" {
		log.Fatal("Telegram token missing, check README")
	}

	b, err := tele.NewBot(tele.Settings{
		Token:     tgToken,
		Poller:    poller,
		Verbose:   verbose,
		ParseMode: tele.ModeHTML,
	})

	if err != nil {
		log.Fatal(err)
	}

	if useWebhook == "" {
		// If we switch from webhook to poller
		// we may need to delete the webhook
		b.RemoveWebhook()
	}

	b.Handle("/help", bot.OnHelp)
	b.Handle("/start", bot.OnHelp)
	b.Handle("/random", bot.OnRandom)
	b.Handle("/randomlang", bot.OnRandomLang)

	b.Handle(tele.OnText, bot.OnHelp)
	//b.Handle(tele.OnQuery, h.OnInlineQuery)

	log.Info("Start bot")
	b.Start()
}
