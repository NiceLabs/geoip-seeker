package shared

type Builder interface {
	AddLocation(location *Location)

	Export() []byte
}
