package callerr

import (
	"errors"
	"fmt"
	"runtime"
)

// New returns an error with caller info.
func New(msg string) error {
	return newCallerErr(errors.New(msg))
}

// Format returns a formatted error with caller info.
func Format(msg string, args ...any) error {
	return newCallerErr(fmt.Errorf(msg, args...))
}

type callerErr struct {
	cause  error
	caller *caller
}

func newCallerErr(err error) *callerErr {
	c := callerErr{cause: err}
	if _, file, line, ok := runtime.Caller(2); ok {
		c.caller = &caller{
			file: file,
			line: line,
		}
	}
	return &c
}

func (e *callerErr) Error() string {
	return fmt.Sprintf("\n[%s] %v", e.caller, e.cause)
}

func (e *callerErr) Cause() error { return e.cause }

func (e *callerErr) Is(err error) bool {
	_, ok := err.(*callerErr)
	return ok
}

type caller struct {
	file string
	line int
}

func (c *caller) String() string {
	if c == nil {
		return "caller unknown"
	}
	return fmt.Sprintf("%s:%d", c.file, c.line)
}
