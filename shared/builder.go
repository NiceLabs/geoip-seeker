package shared

import "io"

type Builder interface {
	AddRecord(location *Record)
	io.WriterTo
}
