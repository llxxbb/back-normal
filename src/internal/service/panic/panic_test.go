package panic

import (
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutOfIndex(t *testing.T) {
	msg := OutOfIndex()
	debug.PrintStack()
	assert.Equal(t, "panic occurred:runtime error: index out of range [3] with length 3", msg)
}

func TestMyPanic(t *testing.T) {
	msg := MyPanic()
	debug.PrintStack()
	assert.Equal(t, "my panic", msg)
}
