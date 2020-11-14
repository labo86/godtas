package dir

type Dir interface {
	Open() error
	Close() error
	Filename(filename string) string
}
