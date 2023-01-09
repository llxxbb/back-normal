package error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_err(t *testing.T) {
	a, err := DoSomeThing("err")
	assert.Empty(t, a)
	assert.Contains(t, err.Error(), "should give 'ok'")
}

func TestError_ok(t *testing.T) {
	a, err := DoSomeThing("ok")
	assert.Equal(t, a, "ok")
	assert.Nil(t, err)
}
