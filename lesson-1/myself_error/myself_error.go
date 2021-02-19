package myself_error

import (
	"fmt"
	"runtime/debug"
	"time"
)


type ErrorWithTraceAndTime struct {
	text string
	trace string
	time   string
}

func New(text string) error {
	return &ErrorWithTraceAndTime{
		text: text,
		trace: string(debug.Stack()),
		time: time.Now().String(),
	}
}

func (e *ErrorWithTraceAndTime) Error() string {
	return fmt.Sprintf("error: %s\ntrace:\n%s\ntime:\n%s", e.text, e.trace, e.time)
}

func StartMyselfError() {
	var err error

	err = New("my error")
	fmt.Println(err)
}

