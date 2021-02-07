package main

import (
	apis "github.com/klenje/wikipediarandombot/pkg/apis"
	log "github.com/sirupsen/logrus"
)

func main() {

	var pages apis.Pages
	howMany := 2
	lang := "en"
	err := apis.GetWiki(&pages, lang, howMany)
	if err != nil {
		log.Error(err)
	}
	log.Warn(len(pages.Items))
	log.Info(pages.Items)
	for p := range pages.Items {
		log.Info(p)
	}
}
