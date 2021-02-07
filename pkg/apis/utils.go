package apis

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
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
