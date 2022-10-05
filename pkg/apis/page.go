package apis

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/Jeffail/gabs/v2"
	log "github.com/sirupsen/logrus"
)

func parseJsonPage(jsonInput []byte) (error, string) {
	jsonParsed, jsonErr := gabs.ParseJSON(jsonInput)
	if jsonErr != nil {
		return jsonErr, ""
	}

	for _, v := range jsonParsed.Path("query.pages").ChildrenMap() {
		val := v.ChildrenMap()["extract"].Data()
		if val == nil {
			return errors.New("missing field extract"), ""
		}
		log.Debug(val)
		return nil, val.(string)
	}
	return nil, ""
}

func getPageFullText(title string, lang string) (error, string) {
	encodedTitle := url.QueryEscape(title)
	uri := fmt.Sprintf("https://%s.wikipedia.org/w/api.php?action=query&format=json&prop=extracts&titles=%s&explaintext=1", lang, encodedTitle)
	body, _, _ := getConnection(uri)
	err, result := parseJsonPage(body)
	if err != nil {
		return err, ""
	}
	return nil, result
}

func ComputeReadingTime(pages *Pages) error {
	for k, page := range pages.Items {
		err, text := getPageFullText(page.Title, page.Pagelanguage)
		if err != nil {
			return err
		}
		words := wordCount(text)
		minutes := timeToRead(words)
		// Update the value in the map
		log.Infof("%s takes %d", page.Title, minutes)
		pages.Items[k].TimeToRead = minutes
	}
	return nil
}
