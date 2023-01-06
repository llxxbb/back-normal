package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoSomething_err(t *testing.T) {
	_, err := DoSomeThing("err")
	assert.Contains(t, err.Error(), "should give 'ok'")
}
