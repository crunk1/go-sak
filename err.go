package sak

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

var (
	defaultErrConfig = errConfig{
		format: "%{filename}:%{line}] %{payload}",
	}
	ErrConfig = defaultErrConfig
)

type errConfig struct {
	format string
}

type sErr struct {
	filename string
	line     int
	payload  string
	msg      string
}

func (e *sErr) Error() string {
	return e.msg
}

func (e *sErr) format() {
	r := strings.NewReplacer("%{filename}", e.filename, "%{line}", strconv.Itoa(e.line), "%{payload}", e.payload)
	e.msg = r.Replace(ErrConfig.format)
}

func Errf(format string, a ...interface{}) error {
	return NewErr(fmt.Sprintf(format, a...))
}

func NewErr(payload interface{}) error {
	e := new(sErr)
	if p, ok := payload.(string); ok {
		e.payload = p
	} else if p, ok := payload.(error); ok {
		e.payload = p.Error()
	} else {
		e.payload = fmt.Sprintf("%#v", payload)
	}

	_, this, _, _ := runtime.Caller(0)
	_, e.filename, e.line, _ = runtime.Caller(1)
	if e.filename == this {
		_, e.filename, e.line, _ = runtime.Caller(2)
	}
	e.format()
	return e
}
