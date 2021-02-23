package apis

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func getConnection(uri string) ([]byte, int, error) {
	client := &http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, 0, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, res.StatusCode, err
	}
	log.Info("HTTP request: " + req.URL.String())
	log.Info("HTTP result: " + res.Status)

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, 0, readErr
	}

	return body, res.StatusCode, nil
}

func wordCount(page string) int {
	words := strings.Fields(page)
	return len(words)
}

func timeToRead(words int) int {
	// Average 250 words per minute
	minutes := words / 250
	// at least 1 minute
	if minutes < 1 {
		return 1
	}
	return minutes
}
