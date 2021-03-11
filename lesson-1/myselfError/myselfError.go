package myselfError

import (
	"fmt"
	"time"
)

type ErrorWithTime struct {
	text string
	time string
}

func StartMyselfError() {
	err := New("my error")
	fmt.Println(err)
}

func New(text string) error {
	return &ErrorWithTime{
		text: text,
		time: time.Now().String(),
	}
}

func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("error: %s\ntime:\n%s", e.text, e.time)
}
