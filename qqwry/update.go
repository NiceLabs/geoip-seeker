package qqwry

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/NiceLabs/geoip-seeker/shared"
)

const (
	updateCopyWrite = "http://update.cz88.net/ip/copywrite.rar"
	updateQQWay     = "http://update.cz88.net/ip/qqwry.rar"
)

type updater struct {
	client             *http.Client
	version, size, key uint32
}

func DownloadUpdate(client *http.Client) (update shared.Update, err error) {
	resp, err := client.Do(newRequest(updateCopyWrite))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	copywrite := new(struct {
		Magic                    [4]byte
		Version, _, Size, _, Key uint32
		Text, Link               [128]byte
	})
	err = binary.Read(resp.Body, binary.LittleEndian, copywrite)
	if err != nil {
		return
	}
	if string(copywrite.Magic[:]) != "CZIP" {
		err = errors.New("magic error")
		return
	}
	update = &updater{
		client:  client,
		version: copywrite.Version,
		size:    copywrite.Size,
		key:     copywrite.Key,
	}
	return
}

func (u *updater) BuildTime() time.Time {
	year, month, day := versionToDate(
		u.version + dateToVersion(1899, 12, 30),
	)

	location := time.FixedZone("CST", +8*3600)
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, location)
}

func (u *updater) Size() uint64 { return uint64(u.size) }

func (u *updater) Download() (payload []byte, err error) {
	resp, err := u.client.Do(newRequest(updateQQWay))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	payload, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return u.decode(payload)
}

func (u *updater) decode(payload []byte) (data []byte, err error) {
	key := u.key
	for index := 0; index < 0x200; index++ {
		key *= 0x805
		key += 1
		key &= 0xFF
		payload[index] = byte(key ^ uint32(payload[index]))
	}
	reader, err := zlib.NewReader(bytes.NewReader(payload))
	if err != nil {
		return
	}
	return ioutil.ReadAll(reader)
}

func newRequest(url string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header = http.Header{
		"Accept":     []string{"text/html, */*"},
		"User-Agent": []string{"Mozilla/3.0 (compatible; Indy Library)"},
	}
	return request
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
