package qqwry

import (
	"net/http"
	"testing"
)

func TestDownloadUpdate(t *testing.T) {
	updater, err := DownloadUpdate(http.DefaultClient)
	if err != nil {
		t.Error(err)
		return
	}
	updater.BuildTime()
	updater.Size()
	_, err = updater.Download()
	if err != nil {
		t.Error(err)
		return
	}
}
