package log

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Logger struct {
	l      *log.Logger
	out    io.Writer
	prefix string
}

func New(out io.Writer, prefix string) *Logger {
	if Reference {
		return &Logger{l: log.New(out, prefix, log.LstdFlags)}
	}
	return &Logger{out: out, prefix: prefix}
}
func (l *Logger) SetOutput(w io.Writer) {
	if Reference {
		l.l.SetOutput(w)
		return
	}
	l.out = w
}
func (l *Logger) SetPrefix(prefix string) {
	if Reference {
		l.l.SetPrefix(prefix)
		return
	}
	l.prefix = prefix
}
func (l *Logger) Output(calldepth int, s string) error {
	if Reference {
		return l.l.Output(calldepth, s)
	}
	var year, month, day = time.Now().Date()
	var _, file, line, _ = runtime.Caller(calldepth)
	var _, err = l.out.Write([]byte(fmt.Sprintf("%s%04d/%02d/%02d %s:%d: %s\n", l.prefix, year, month, day, file, line, s)))
	return err
}
