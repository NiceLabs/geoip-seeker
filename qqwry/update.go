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
	updateCopyWrite = "http://update.cz88.net/ip/copywrite.rar"
	updateQQWay     = "http://update.cz88.net/ip/qqwry.rar"
)

type Update struct {
	version uint32
	size    uint32
	key     uint32
}

func DownloadUpdate() (*Update, error) {
	resp, err := http.Get(updateCopyWrite)
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
		Text    [128]byte
		Link    [128]byte
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

func (update *Update) BuildTime() time.Time {
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

// see https://github.com/shuax/LocateIP/blob/master/loci/cz_update.c#L23-L29
func dateToVersion(year, month, day uint32) uint32 {
	month = (month + 9) % 12
	year = year - month/10
	day = 365*year + year/4 - year/100 + year/400 + (month*153+2)/5 + day - 1
	return day
}

// see https://github.com/shuax/LocateIP/blob/master/loci/cz_update.c#L31-L41
func versionToDate(version uint32) (year, month, day int) {
	y := (version*33 + 999) / 12053
	t := version - y*365 - y/4 + y/100 - y/400
	m := (t*5+2)/153 + 2

	year = int(y + m/12)
	month = int(m%12 + 1)
	day = int(t - (m*153-304)/5 + 1)
	return
}
