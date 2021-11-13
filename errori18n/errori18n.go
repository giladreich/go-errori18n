package errori18n

import (
	"fmt"
)

type Error struct {
	ErrorCode ErrorCode
	args      []interface{}
}

func (e *Error) String() string {
	return fmt.Sprintf(e.ErrorCode.String(), e.args...)
}

func Errorf(code ErrorCode, args ...interface{}) *Error {
	return &Error{
		ErrorCode: code,
		args:      args,
	}
}
