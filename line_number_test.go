package callerr

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nextLineNumber(t *testing.T) {
	line := nextLineNumber(t)
	require.Equal(t, 12, line)
}

func nextLineNumber(t *testing.T) int {
	_, _, line, ok := runtime.Caller(1)
	require.True(t, ok)
	return line + 1
}
