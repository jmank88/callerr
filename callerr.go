package callerr

import (
	"fmt"
	"io"
	"path/filepath"
	"runtime"
)

// New returns an error with caller info.
func New(msg string) error {
	return fmt.Errorf("\n%s "+msg, getCaller())
}

// Format returns a formatted error with caller info via fmt.Errorf.
func Format(msg string, args ...any) error {
	args = append([]any{getCaller()}, args...)
	return fmt.Errorf("\n%s "+msg, args...)
}

// Println prints a formatted log with caller info via fmt.Println.
func Println(args ...any) {
	fmt.Println(append([]any{fmt.Sprintf("\n%s", getCaller())}, args...)...)
}

// Fprintln writes a formatted log with caller info via fmt.Fprintln.
func Fprintln(w io.Writer, args ...any) (int, error) {
	return fmt.Fprintln(w, append([]any{fmt.Sprintf("\n%s", getCaller())}, args...)...)
}

// Sprintln returns a formatted log with caller info via fmt.Sprintln.
func Sprintln(w io.Writer, args ...any) string {
	return fmt.Sprintln(append([]any{fmt.Sprintf("\n%s", getCaller())}, args...)...)
}

// Printf prints a formatted log with caller info via fmt.Printf.
func Printf(msg string, args ...any) {
	args = append([]any{getCaller()}, args...)
	fmt.Printf("\n%s "+msg, args...)
}

// Fprintf writes a formatted log with caller info via fmt.Fprintf.
func Fprintf(w io.Writer, msg string, args ...any) (n int, err error) {
	args = append([]any{getCaller()}, args...)
	return fmt.Fprintf(w, "\n%s "+msg, args...)
}

// Sprintf returns a formatted log with caller info via fmt.Sprintf.
func Sprintf(msg string, args ...any) string {
	args = append([]any{getCaller()}, args...)
	return fmt.Sprintf("\n%s "+msg, args...)
}

// getCaller returns an arg to be formatted. Either a string or a fmt.Stringer.
func getCaller() any {
	if pc, file, line, ok := runtime.Caller(2); ok {
		return caller{pc, file, line}
	}
	return "caller unknown"
}

// caller holds runtime.Caller data and implements fmt.Stringer.
type caller struct {
	pc   uintptr
	file string
	line int
}

// String returns formatted caller info, including file, line number, and (if available) func name.
func (c caller) String() string {
	frames := runtime.CallersFrames([]uintptr{c.pc})
	f, _ := frames.Next()
	if f == (runtime.Frame{}) { // no frame available for func name
		return fmt.Sprintf("\t%s:%d", c.file, c.line)
	}
	return fmt.Sprintf("%s()\n\t%s:%d", filepath.Base(f.Function), c.file, c.line)
}
