package log

type Options struct {
	level int
}

const (
	NULL = iota
	TRACE = 1
	DEBUG = 2
	INFO = 3
	WARNING = 4
	ERROR = 5
	FATAL = 6
)