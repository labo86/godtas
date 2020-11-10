package dir

type Interface interface {
	Init() error
	Close() error
	Filename(filename string) string
}
