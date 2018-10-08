package shared

import "io"

type Builder interface {
	AddRecord(record *Record)
	io.WriterTo
}
