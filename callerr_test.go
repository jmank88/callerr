package callerr

import (
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
		// drop pkg/callerErr/callerr_test.go
		file = filepath.Dir(file) // drop callerr_test.go
		file = filepath.Dir(file) // drop callerr
		basepath = file
	}
}

func TestNew(t *testing.T) {
	err := New("test")
	assert.Equal(t, fmt.Sprintf(`
[%s/callerr/callerr_test.go:26] test`, basepath), err.Error())
	assert.ErrorIs(t, err, &callerErr{})
}

func TestFormat(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		bazLine := nextLineNumber(t)
		bazErr := New("baz")
		line := nextLineNumber(t)
		err := Format("foo: %w", Format("bar: %s", bazErr))
		got := err.Error()
		assert.Equal(t, got, fmt.Sprintf(`
[%[1]s/callerr/callerr_test.go:%[2]d] foo: 
[%[1]s/callerr/callerr_test.go:%[2]d] bar: 
[%[1]s/callerr/callerr_test.go:%[3]d] baz`, basepath, line, bazLine))

		var ce *callerErr
		require.ErrorAs(t, err, &ce)
		assert.ErrorIs(t, ce.cause, bazErr)
		assert.ErrorIs(t, err, bazErr)
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
		assert.Equal(t, got, fmt.Sprintf(`
[%[1]s/callerr/callerr_test.go:%[2]d] foo: 
[%[1]s/callerr/callerr_test.go:%[2]d] this
is
a
multi-line err: 
[%[1]s/callerr/callerr_test.go:%[3]d] baz`, basepath, line, bazLine))
	})
}
