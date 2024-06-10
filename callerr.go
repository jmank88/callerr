package callerr

import (
	"errors"
	"fmt"
	"runtime"
)

type callerErr struct {
	cause  error
	caller string
}

func newcallerErr(err error) *callerErr {
	var caller string
	_, file, line, ok := runtime.Caller(2)
	if ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	} else {
		caller = "caller unknown"
	}
	return &callerErr{
		cause:  err,
		caller: caller,
	}
}

func (e *callerErr) Error() string {
	return fmt.Sprintf("\n[%s] %v", e.caller, e.cause)
}

func (e *callerErr) Cause() error { return e.cause }

func (e *callerErr) Is(err error) bool {
	_, ok := err.(*callerErr)
	return ok
}

func New(msg string) error {
	return newcallerErr(errors.New(msg))
}

func Format(msg string, args ...any) error {
	return newcallerErr(fmt.Errorf(msg, args...))
}
