package callerr

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var basepath string

func init() {
	_, file, _, ok := runtime.Caller(0)
	if ok {
		// drop callerr/callerr_test.go
		file = filepath.Dir(file) // drop callerr_test.go`
		file = filepath.Dir(file) // drop callerr
		basepath = file
	}
}

func TestNew(t *testing.T) {
	line := nextLineNumber(t)
	err := New("test")
	assert.ErrorContains(t, err, fmt.Sprintf(`
callerr.TestNew()
	%s/callerr/callerr_test.go:%d test`, basepath, line))
}

func TestFormat(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		bazLine := nextLineNumber(t)
		bazErr := New("baz")
		line := nextLineNumber(t)
		err := Format("foo: %w", Format("bar: %w", bazErr))
		got := err.Error()

		assert.Equal(t, fmt.Sprintf(`
callerr.TestFormat.func1()
	%[1]s/callerr/callerr_test.go:%[2]d foo: 
callerr.TestFormat.func1()
	%[1]s/callerr/callerr_test.go:%[2]d bar: 
callerr.TestFormat.func1()
	%[1]s/callerr/callerr_test.go:%[3]d baz`, basepath, line, bazLine), got)

		cause := errors.Unwrap(err)
		require.NotNil(t, cause)
		assert.ErrorIs(t, cause, bazErr)
		assert.ErrorIs(t, err, bazErr)
	})

	t.Run("wrapped", func(t *testing.T) {
		bazLine := nextLineNumber(t)
		bazErr := New("baz")
		line := nextLineNumber(t)
		err := Format("foo: %w", Format("this error message is only a single line but it is very long so that it ends up being wrapped, hopefully even more than once: %w", bazErr))
		got := err.Error()
		msg := fmt.Sprintf(`
callerr.TestFormat.func2()
	%[1]s/callerr/callerr_test.go:%[2]d foo: 
callerr.TestFormat.func2()
	%[1]s/callerr/callerr_test.go:%[2]d this error message is only a single line but it is very long so that it ends up being wrapped, hopefully even more than once: 
callerr.TestFormat.func2()
	%[1]s/callerr/callerr_test.go:%[3]d baz`, basepath, line, bazLine)
		t.Log("Expect:", msg)
		assert.Equal(t, msg, got)
	})

	t.Run("multiline", func(t *testing.T) {
		bazLine := nextLineNumber(t)
		bazErr := New("baz")
		line := nextLineNumber(t)
		err := Format("foo: %w", Format(`this
is
a
multi-line err: %s`, bazErr))
		got := err.Error()
		assert.Equal(t, fmt.Sprintf(`
callerr.TestFormat.func3()
	%[1]s/callerr/callerr_test.go:%[2]d foo: 
callerr.TestFormat.func3()
	%[1]s/callerr/callerr_test.go:%[2]d this
is
a
multi-line err: 
callerr.TestFormat.func3()
	%[1]s/callerr/callerr_test.go:%[3]d baz`, basepath, line, bazLine), got)
	})

	//TODO only bottom of stack
	//TODO only top of stack
}

func Benchmark_getCaller(b *testing.B) {
	b.Run("inline", func(b *testing.B) {
		var arg any
		for b.Loop() {
			arg = getCallerInline()
		}
		_ = arg
	})
	b.Run("deferred", func(b *testing.B) {
		var arg any
		for b.Loop() {
			arg = getCaller()
		}
		_ = arg
	})
}

// getCallerInline is an inline version of getCaller.
func getCallerInline() string {
	if pc, file, line, ok := runtime.Caller(2); ok {
		frames := runtime.CallersFrames([]uintptr{pc})
		f, _ := frames.Next()
		if f == (runtime.Frame{}) { // no frame available for func name
			return fmt.Sprintf("\t%s:%d", file, line)
		}
		return fmt.Sprintf("%s()\n\t%s:%d", filepath.Base(f.Function), file, line)
	}
	return "caller unknown"
}
