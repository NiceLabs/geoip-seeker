package ipip_net

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	updateAPI = "http://api.ipip.net/api.php?a=ipdb"
)

type Update struct {
	version string
	url     string
}

func DownloadUpdate() (*Update, error) {
	resp, err := http.Get(updateAPI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fields := strings.Split(string(data), "|")
	if len(fields) != 3 {
		return nil, errors.New("request error")
	}
	update := new(Update)
	update.version = fields[1]
	update.url = fields[2]
	return update, nil
}

func (update *Update) PublishDate() time.Time {
	return resolvePublishDate(update.version)
}

func (update *Update) DownloadDB() ([]byte, error) {
	resp, err := http.Get(update.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
