package qqwry

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	updateCopywrite = "http://update.cz88.net/ip/copywrite.rar"
	updateQQWay     = "http://update.cz88.net/ip/qqwry.rar"
)

type Update struct {
	version uint32
	size    uint32
	key     uint32
}

func DownloadUpdate() (*Update, error) {
	resp, err := http.Get(updateCopywrite)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	copywrite := new(struct {
		Magic   [4]byte
		Version uint32
		_       uint32
		Size    uint32
		_       uint32
		Key     uint32
	})
	if err := binary.Read(resp.Body, binary.LittleEndian, copywrite); err != nil {
		return nil, err
	}
	if string(copywrite.Magic[:]) != "CZIP" {
		return nil, errors.New("magic error")
	}
	update := new(Update)
	update.version = copywrite.Version
	update.size = copywrite.Size
	update.key = copywrite.Key
	return update, nil
}

func (update *Update) decodeQQWay(data []byte) []byte {
	key := update.key
	for index := 0; index < 0x200; index++ {
		key *= 0x805
		key += 1
		key &= 0xFF
		data[index] = byte(key ^ uint32(data[index]))
	}
	reader, _ := zlib.NewReader(bytes.NewReader(data))
	data, _ = ioutil.ReadAll(reader)
	return data
}

func (update *Update) PublishDate() time.Time {
	year, month, day := versionToDate(
		update.version + dateToVersion(1899, 12, 30),
	)

	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, location)
}

func (update *Update) Size() uint32 {
	return update.size
}

func (update *Update) DownloadDB() ([]byte, error) {
	resp, err := http.Get(updateQQWay)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return update.decodeQQWay(payload), nil
}
