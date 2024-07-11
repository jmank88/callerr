package callerr

import (
	"fmt"
	"runtime"
)

// New returns an error with caller info.
func New(msg string) error {
	return fmt.Errorf("\n[%s] "+msg, getCaller())
}

// Format returns a formatted error with caller info.
func Format(msg string, args ...any) error {
	args = append([]any{getCaller()}, args...)
	return fmt.Errorf("\n[%s] "+msg, args...)
}

func getCaller() string {
	if _, file, line, ok := runtime.Caller(2); ok {
		return fmt.Sprintf("%s:%d", file, line)
	}
	return "caller unknown"
}
