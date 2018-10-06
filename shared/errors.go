package shared

import "errors"

var (
	ErrFileSize          = errors.New("ip database file size error")
	ErrModeError         = errors.New("ip database mode error")
	ErrMetaData          = errors.New("ip database metadata error")
	ErrDatabaseError     = errors.New("database error")
	ErrIPFormat          = errors.New("query ip format error")
	ErrNoSupportLanguage = errors.New("language not support")
	ErrNoSupportIPv4     = errors.New("ipv4 not support")
	ErrNoSupportIPv6     = errors.New("ipv6 not support")
	ErrInvalidIPv4       = errors.New("invalid ipv4 address")
	ErrDataNotExists     = errors.New("data is not exists")
	ErrIndexOutOfRange   = errors.New("index out of range")
)
